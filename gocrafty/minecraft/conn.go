package minecraft

import (
	"bufio"
	"errors"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/handshake"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/login"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/status"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
	"io"
	"net"
	"sync"
)

type Conn struct {
	listener *Listener

	sendMutex sync.Mutex
	conn      net.Conn

	reader *bufio.Reader

	// The buffer used for recving data
	recvBuffer   [65565]byte
	decompBuffer [65565]byte // TODO: implement compression/decompression and give sense to this buffer cuz it's useless rn xd

	State int32
}

func NewConn(listener *Listener, conn net.Conn) *Conn {
	return &Conn{
		listener: listener,
		conn:     conn,
		reader:   bufio.NewReader(conn),
		State:    types.StateHandshaking,
	}
}

func (c *Conn) WritePacket(p packet.Packet) error {
	packetWriter := &protocol.Writer{}
	packetWriter.VarInt(p.ID())
	p.Marshal(packetWriter)

	payload := packetWriter.Bytes()

	sendPacket := &protocol.Writer{}
	sendPacket.VarInt(int32(len(payload)))
	sendPacket.WriteBytes(payload)

	c.sendMutex.Lock()
	defer c.sendMutex.Unlock()

	_, err := c.conn.Write(sendPacket.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) ReadPacket() (packet.Packet, error) {
	packetLength, err := c.varInt()
	if err != nil {
		return nil, err
	}

	_, err = io.ReadFull(c.reader, c.recvBuffer[:packetLength])
	if err != nil {
		return nil, err
	}

	reader := protocol.Reader{Data: c.recvBuffer[:packetLength]}
	pool := packet.NewPool()

	var packetId int32
	reader.VarInt(&packetId)

	p, ok := pool.Get(c.State, packetId)
	if !ok {
		return nil, err
	}

	p.Unmarshal(&reader)

	c.handlePacket(p)

	return p, nil
}

func (c *Conn) handlePacket(p packet.Packet) {
	c.listener.logger.Debugf("Received packet with ID %d", p.ID())

	switch c.State {

	// Handshaking
	case types.StateHandshaking:
		switch dp := p.(type) {
		case *handshake.Handshake:
			switch dp.NextState {
			case 1:
				c.State = types.StateStatus

			case 2:
				if dp.ProtocolVersion == ProtocolVersion {
					c.State = types.StateLogin

					if c.listener.playerCount.Load() >= int32(c.listener.maxPlayers) {
						c.WritePacket(&login.Disconnect{
							Reason: &types.Chat{
								Text:  "Server is full",
								Bold:  true,
								Color: "red",
							},
						})

						return
					}

					c.listener.playerCount.Add(1)
					defer c.listener.playerCount.Add(-1)

					// TODO: Send login success packet
					c.WritePacket(&login.Disconnect{
						Reason: &types.Chat{
							Text: "Not implemented yet",
						},
					})
				} else {
					c.WritePacket(&login.Disconnect{
						Reason: &types.Chat{
							Text:  "Invalid protocol version",
							Bold:  true,
							Color: "red",
						},
					})
				}
			}
		}

	// Status
	case types.StateStatus:
		switch p.(type) {
		case *status.StatusRequest:
			// TODO: Send real data instead of hardcoded one
			c.WritePacket(&status.StatusResponse{
				JSONResponse: &status.StatusResponseData{
					Version: &status.StatusResponseDataVersion{
						Name:     Version,
						Protocol: ProtocolVersion,
					},
					Players: &status.StatusResponseDataPlayers{
						Max:    int32(c.listener.maxPlayers),
						Online: c.listener.playerCount.Load(),
					},
					Description: &status.StatusResponseDataDescription{
						Text: c.listener.name,
					},
				},
			})

		case *status.PingRequest:
			c.WritePacket(&status.PingResponse{
				PingTime: p.(*status.PingRequest).PingTime,
			})
		}
	}
}

func (c *Conn) Close() error {
	// TODO: Write disconnect packet

	return c.conn.Close()
}

func (c *Conn) varInt() (int, error) {
	numRead := 0
	result := 0
	for {
		read, err := c.reader.ReadByte()
		if err != nil {
			return 0, err
		}
		value := read & 0b01111111
		result |= int(value) << (7 * numRead)

		numRead++
		if numRead > 5 {
			return 0, errors.New("varint is too big")
		}

		if (read & 0b10000000) == 0 {
			return result, nil
		}
	}
}
