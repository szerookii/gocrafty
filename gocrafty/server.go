package gocrafty

import (
	"errors"
	"github.com/szerookii/gocrafty/gocrafty/minecraft"
	"sync/atomic"
)

type Server struct {
	// Config is the configuration of the server.
	config *ServerConfig
	// started is whether the server has been started or not.
	started atomic.Bool
	// listener is the listener of the server.
	listener *minecraft.Listener
	// conns is a map of all the connections to the server.
	conns map[*minecraft.Conn]struct{}
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		config: config,
		conns:  make(map[*minecraft.Conn]struct{}),
	}
}

func (s *Server) Listen() error {
	if s.started.Load() {
		return errors.New("server already started")
	}

	s.listener = minecraft.NewListener(s.config.Logger, s.config.ServerName, s.config.BoundAddress, s.config.MaxPlayers)

	_, err := s.listener.Listen()

	if err != nil {
		return err
	}

	s.started.Store(true)

	// TODO: ???

	return nil
}
