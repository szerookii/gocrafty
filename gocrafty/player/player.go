package player

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/socket"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
	"github.com/szerookii/gocrafty/gocrafty/world"
	"sync"
)

type Player struct {
	sync.RWMutex

	// Entity fields
	eid        int32
	gamemode   GameMode
	dimension  types.Dimension
	difficulty Difficulty
	world      *world.World

	logger logger.Logger

	username string
	uuid     uuid.UUID

	conn *socket.Conn
}

func New(logger logger.Logger, username string, uuid uuid.UUID, conn *socket.Conn) *Player {
	return &Player{
		logger: logger,

		eid:        world.NextEID(),
		gamemode:   GameModeSurvival,
		dimension:  types.Overworld,
		difficulty: DifficultyNormal,

		username: username,
		uuid:     uuid,

		conn: conn,
	}
}

func (p *Player) JoinGame() {
	p.conn.WritePacket(&play.JoinGame{
		EID:              p.EID(),
		Gamemode:         uint8(p.Gamemode()),
		Dimension:        byte(p.Dimension()),
		Difficulty:       uint8(p.Difficulty()),
		MaxPlayers:       20,
		LevelType:        p.Dimension().LevelType(),
		ReducedDebugInfo: true,
	})

	// TODO: send chunks
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

func (p *Player) EID() int32 {
	p.RLock()
	defer p.RUnlock()

	return p.eid
}

func (p *Player) SetEID(eid int32) {
	p.Lock()
	defer p.Unlock()

	p.eid = eid
}

func (p *Player) Gamemode() GameMode {
	p.RLock()
	defer p.RUnlock()

	return p.gamemode
}

func (p *Player) SetGamemode(gamemode GameMode) {
	p.Lock()
	defer p.Unlock()

	p.gamemode = gamemode

	// TODO: Send gamemode change packet to the client.
}

func (p *Player) Dimension() types.Dimension {
	p.RLock()
	defer p.RUnlock()

	return p.dimension
}

func (p *Player) SetDimension(dimension types.Dimension) {
	p.Lock()
	defer p.Unlock()

	p.dimension = dimension

	// TODO: Send dimension change packet to the client.
}

func (p *Player) Difficulty() Difficulty {
	p.RLock()
	defer p.RUnlock()

	return p.difficulty
}

func (p *Player) SetDifficulty(difficulty Difficulty) {
	p.Lock()
	defer p.Unlock()

	p.difficulty = difficulty

	// TODO: Send difficulty change packet to the client.
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
