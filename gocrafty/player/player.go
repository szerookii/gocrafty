package player

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/socket"
	"sync"
)

type Player struct {
	sync.RWMutex

	logger logger.Logger

	username string
	uuid     uuid.UUID

	conn *socket.Conn
}

func New(logger logger.Logger, username string, uuid uuid.UUID, session *socket.Conn) *Player {
	return &Player{
		logger: logger,

		username: username,
		uuid:     uuid,
		conn:     session,
	}
}

func (p *Player) Username() string {
	p.RLock()
	defer p.RUnlock()

	return p.username
}

func (p *Player) UUID() uuid.UUID {
	p.RLock()
	defer p.RUnlock()

	return p.uuid
}

func (p *Player) Conn() *socket.Conn {
	p.RLock()
	defer p.RUnlock()

	return p.conn
}

func (p *Player) Disconnect(reason string) {
	p.Conn().Close(reason)
}

func (p *Player) Logger() logger.Logger {
	p.RLock()
	defer p.RUnlock()

	return p.logger
}