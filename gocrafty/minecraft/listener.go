package minecraft

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/handler"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/handshake"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/login"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/status"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/socket"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
	"github.com/szerookii/gocrafty/gocrafty/player"
	"io"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Listener struct {
	sync.RWMutex

	logger     logger.Logger
	name       string
	favicon    string
	address    string
	maxPlayers int
	onlineMode bool

	PlayerCount atomic.Int32
	listener    net.Listener
	pool        packet.Pool

	playerList *player.PlayerList

	incoming   chan *player.Player
	disconnect chan *player.Player

	close chan struct{}
}

func NewListener(incoming, disconnect chan *player.Player, logger logger.Logger, name, favicon, addr string, maxPlayers int, onlineMode bool) *Listener {
	return &Listener{
		logger:     logger,
		name:       name,
		favicon:    favicon,
		address:    addr,
		maxPlayers: maxPlayers,
		onlineMode: onlineMode,

		pool: packet.NewPool(),

		playerList: player.NewPlayerList(),

		incoming:   incoming,
		disconnect: disconnect,

		close: make(chan struct{}),
	}
}

func (l *Listener) Listen() (*Listener, error) {
	listener, err := net.Listen("tcp", l.address)

	if err != nil {
		return nil, errors.New("failed to listen on address " + l.address + ": " + err.Error())
	}

	l.listener = listener

	go l.listen()

	return l, nil
}

func (l *Listener) listen() {
	for {
		netConn, err := l.listener.Accept()
		if err != nil {
			l.logger.Errorf("Got an error accepting connection: %v", err)
			continue
		}

		l.createConn(netConn)
	}
}
	}()

	defer func() {
		// TODO: incoming channel
		close(l.close)
		_ = l.Close()
	}()

	for {
		netConn, err := l.listener.Accept()
		if err != nil {
			return
		}

		l.createConn(netConn)
	}
}

func (l *Listener) createConn(netConn net.Conn) {
	conn := socket.NewConn(netConn)

	go l.handleConn(conn)
}

func (l *Listener) handleConn(conn *socket.Conn) {
	defer func() {
		_ = conn.Close("")

		if r := recover(); r != nil {
			l.logger.Debugf("panic in connection handler: %v\n", string(debug.Stack()))
		}
	}()

	for {
		p, err := conn.ReadPacket()

		if err != nil {
			if p, ok := l.playerList.Get(conn.UUID()); ok {
				l.logger.Infof("Player %s disconnected", conn.Username())

				l.PlayerCount.Add(-1)
				l.RemovePlayer(p)
			}

			if !errors.Is(err, io.EOF) {
				l.logger.Errorf("Got an error receving packet: %v", err)
			}

			break
		}

		if conn.State() == types.StateDisconnect {
			break
		}

		if conn.State() == types.StatePlay {
			pl, ok := l.playerList.Get(conn.UUID())
			if !ok {
				l.logger.Errorf("Got a packet from a player that is not in the player list")
				break
			}

			handler, ok := handler.GetHandler(p.ID())
			if !ok {
				l.logger.Errorf("Got a packet but no handler for it: 0x%x", p.ID())
				continue
			}

			handler.Handle(pl, p)
		} else {
			l.handlePacket(conn, p)
		}

		select {
		case <-l.close:
			return

		default:
			// loop
		}
	}
}

func (l *Listener) handlePacket(c *socket.Conn, p packet.Packet) {
	switch c.State() {

	// Handshaking
	case types.StateHandshaking:
		switch dp := p.(type) {
		case *handshake.Handshake:
			switch dp.NextState {
			case 1:
				c.SetState(types.StateStatus)

			case 2:
				if dp.ProtocolVersion == ProtocolVersion {
					if l.PlayerCount.Load() >= int32(l.MaxPlayers()) {
						c.WritePacket(&login.Disconnect{
							Reason: types.Chat{
								Text:  "Server is full",
								Bold:  true,
								Color: "red",
							},
						})

						c.SetState(types.StateDisconnect)

						return
					}

					l.PlayerCount.Add(1)

					c.SetState(types.StateLogin)
				} else {
					c.WritePacket(&login.Disconnect{
						Reason: types.Chat{
							Text:  fmt.Sprintf("Outdated server! I'm still on %s", Version),
							Bold:  true,
							Color: "red",
						},
					})

					c.SetState(types.StateDisconnect)
				}
			}
		}

	// Status
	case types.StateStatus:
		switch dp := p.(type) {
		case *status.StatusRequest:
			c.WritePacket(&status.StatusResponse{
				JSONResponse: &status.StatusResponseData{
					Version: &status.StatusResponseDataVersion{
						Name:     Version,
						Protocol: ProtocolVersion,
					},
					Players: &status.StatusResponseDataPlayers{
						Max:    int32(l.MaxPlayers()),
						Online: l.PlayerCount.Load(),
					},
					Description: &status.StatusResponseDataDescription{
						Text: l.Name(),
					},
					Favicon: l.Favicon(),
				},
			})

		case *status.PingRequest:
			c.WritePacket(&status.PingResponse{
				PingTime: dp.PingTime,
			})
		}

	// Login
	case types.StateLogin:
		switch dp := p.(type) {
		case *login.LoginStart:
			l.Logger().Infof("Player %s connected", dp.Username)
			c.SetState(types.StatePlay)

			// TODO: Verify player online mode

			// TODO: Support compression
			c.WritePacket(&login.SetCompression{
				Threshold: -1,
			})

			if l.OnlineMode() {
				// TODO: Send encryption request
			} else {
				uuid := uuid.New()
				uuid.UnmarshalText([]byte("OfflinePlayer:" + dp.Username))

				c.SetUsername(dp.Username)
				c.SetUUID(uuid)

				c.WritePacket(&login.LoginSuccess{
					UUID:     uuid.String(),
					Username: dp.Username,
				})

				l.AddPlayer(c)
			}
		}
	}
}

func (l *Listener) AddPlayer(c *socket.Conn) {
	p := player.New(l.logger, l.playerList, c.Username(), c.UUID(), c)

	if _, ok := l.playerList.Get(c.UUID()); ok {
		p.Disconnect("You are already logged in!")

		return
	}

	go func() {
		l.incoming <- p
	}()

	go p.Conn().KeepAliveTicker()

	p.JoinGame()
}

func (l *Listener) RemovePlayer(p *player.Player) {
	p.LeaveGame()

	go func() {
		l.disconnect <- p
	}()
}

func (l *Listener) Close() error {
	return l.listener.Close()
}

func (l *Listener) Pool() packet.Pool {
	l.RLock()
	defer l.RUnlock()

	return l.pool
}

func (l *Listener) Name() string {
	l.RLock()
	defer l.RUnlock()

	return l.name
}

func (l *Listener) Favicon() string {
	l.RLock()
	defer l.RUnlock()

	return l.favicon
}

func (l *Listener) Address() string {
	l.RLock()
	defer l.RUnlock()

	return l.address
}

func (l *Listener) MaxPlayers() int {
	l.RLock()
	defer l.RUnlock()

	return l.maxPlayers
}

func (l *Listener) OnlineMode() bool {
	l.RLock()
	defer l.RUnlock()

	return l.onlineMode
}

func (l *Listener) Logger() logger.Logger {
	l.RLock()
	defer l.RUnlock()

	return l.logger
}
