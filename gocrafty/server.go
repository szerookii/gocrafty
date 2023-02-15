package gocrafty

import (
	"errors"
	"github.com/szerookii/gocrafty/gocrafty/minecraft"
	"github.com/szerookii/gocrafty/gocrafty/player"
	"sync/atomic"
)

type Server struct {
	config   *ServerConfig
	started  atomic.Bool
	listener *minecraft.Listener

	incoming   chan *player.Player
	disconnect chan *player.Player
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		config:     config,
		incoming:   make(chan *player.Player),
		disconnect: make(chan *player.Player), // TODO: maybe its useless
	}
}

func (s *Server) Listen() error {
	if s.started.Load() {
		return errors.New("server already started")
	}

	s.listener = minecraft.NewListener(s.incoming, s.disconnect, s.config.Logger, s.config.ServerName, s.config.BoundAddress, s.config.MaxPlayers, s.config.OnlineMode)

	_, err := s.listener.Listen()

	if err != nil {
		return err
	}

	s.started.Store(true)

	return nil
}

func (s *Server) Accept() *player.Player {
	if !s.started.Load() {
		return nil
	}

	return <-s.incoming
}

func (s *Server) Listener() *minecraft.Listener {
	return s.listener
}
