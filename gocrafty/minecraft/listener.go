package minecraft

import (
	"errors"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"io"
	"net"
	"sync/atomic"
	"time"
)

type Listener struct {
	logger     logger.Logger
	name       string
	address    string
	maxPlayers int

	playerCount atomic.Int32
	listener    net.Listener
	pool        packet.Pool

	close chan struct{}
}

func NewListener(logger logger.Logger, name, addr string, maxPlayers int) *Listener {
	return &Listener{
		logger:     logger,
		name:       name,
		address:    addr,
		maxPlayers: maxPlayers,

		pool:  packet.NewPool(),
		close: make(chan struct{}),
	}
}

func (l *Listener) Listen() (*Listener, error) {
	listener, err := net.Listen("tcp4", l.address)

	if err != nil {
		return nil, errors.New("failed to listen on address " + l.address + ": " + err.Error())
	}

	l.listener = listener

	go l.listen()

	return l, nil
}

func (l *Listener) listen() {
	go func() {
		ticker := time.NewTicker(time.Second * 4)
		defer ticker.Stop()
		for {
			select {
			// TODO: update data packet

			case <-l.close:
				return
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
			// The underlying listener was closed, meaning we should return immediately so this listener can
			// close too.
			return
		}

		l.createConn(netConn)
	}
}

func (l *Listener) createConn(netConn net.Conn) {
	conn := NewConn(netConn)

	if l.playerCount.Load() >= int32(l.maxPlayers) {
		// TODO: kick player
		conn.Close()
		return
	}

	l.playerCount.Add(1)

	go l.handleConn(conn)
}

func (l *Listener) handleConn(conn *Conn) {
	defer func() {
		conn.Close()
		l.playerCount.Add(-1)

		if r := recover(); r != nil {
			l.logger.Debugf("panic in connection handler: %v", r)
		}
	}()

	for {
		_, err := conn.ReadPacket()

		if err != nil {
			if err == io.EOF && conn.State == StatePlay {
				l.logger.Errorf("error reading packet: %v", err)
			}

			return
		}

		// TODO: Handle
	}
}

func (l *Listener) Close() error {
	return l.listener.Close()
}
