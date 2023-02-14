package minecraft

import (
	"bufio"
	"bytes"
	"github.com/kataras/golog"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"io"
	"net"
	"sync"
)

const (
	StateHandshaking = iota
	StateStatus
	StateLogin
	StatePlay
)

type Conn struct {
	sendMutex sync.Mutex
	conn      net.Conn
	reader    *protocol.Reader

	// The buffer used for recving data
	recvBuffer   [65565]byte
	decompBuffer [65565]byte

	State int
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		reader: protocol.NewReader(bufio.NewReader(conn)),
		State:  StateHandshaking,
	}
}

func (c *Conn) WritePacket(p packet.Packet) error {

	return nil
}

func (c *Conn) ReadPacket() (packet.Packet, error) {
	var packetLength int32
	c.reader.VarInt(&packetLength)

	_, err := io.ReadFull(c.reader, c.recvBuffer[:packetLength])
	if err != nil {
		return nil, err
	}

	reader := protocol.NewReader(bytes.NewReader(c.recvBuffer[:packetLength]))
	pool := packet.NewPool()

	var packetId int32
	reader.VarInt(&packetId)

	p, ok := pool.Get(c.State != StateHandshaking, packetId)
	if !ok {
		return nil, err
	}

	p.Unmarshal(reader)

	golog.Println(p)

	return p, nil
}

func (c *Conn) Close() error {
	// TODO: Write disconnect packet

	return c.conn.Close()
}
