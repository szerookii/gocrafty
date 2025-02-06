package player

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/socket"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
	"github.com/szerookii/gocrafty/gocrafty/world"
	"github.com/szerookii/gocrafty/gocrafty/world/chunk"
	"math"
	"sync"
	"sync/atomic"
)

type Player struct {
	sync.RWMutex

	logger     logger.Logger
	playerList *PlayerList

	username string
	uuid     uuid.UUID

	// Entity fields
	Eid        atomic.Int32
	gamemode   GameMode
	dimension  types.Dimension
	difficulty Difficulty
	world      *world.World
	position   mgl64.Vec3
	yaw        float32
	pitch      float32
	OnGround   atomic.Bool
	sneaking   atomic.Bool
	Sprinting  atomic.Bool

	conn *socket.Conn
}

func New(logger logger.Logger, playerList *PlayerList, username string, uuid uuid.UUID, conn *socket.Conn) *Player {
	p := &Player{
		logger:     logger,
		playerList: playerList,

		gamemode:   GameModeSurvival,
		dimension:  types.Overworld,
		difficulty: DifficultyNormal,
		position:   mgl64.Vec3{0, 50, 0},

		username: username,
		uuid:     uuid,

		conn: conn,
	}

	p.Eid.Store(world.NextEID())
	p.OnGround.Store(true)

	return p
}

func (p *Player) JoinGame() {
	p.PlayerList().Add(p)

	p.Conn().WritePacket(&play.JoinGame{
		EID:              p.Eid.Load(),
		Gamemode:         uint8(p.Gamemode()),
		Dimension:        byte(p.Dimension()),
		Difficulty:       uint8(p.Difficulty()),
		MaxPlayers:       20,
		LevelType:        p.Dimension().LevelType(),
		ReducedDebugInfo: true,
	})

	p.Conn().WritePacket(&play.ClientPlayerPositionAndLook{
		X:     p.Position().X(),
		Y:     p.Position().Y(),
		Z:     p.Position().Z(),
		Yaw:   p.Yaw(),
		Pitch: p.Pitch(),
		Flags: 0,
	})

	// TODO: Save world and player data to the database.
	for x := int32(-5); x <= 1; x++ {
		for z := int32(-5); z <= 1; z++ {
			p.Conn().WritePacket(chunk.GenerateFlatChunk(x, z).Marshal())

			p.logger.Infof("Sent chunk %d %d", x, z)
		}
	}

	// TODO: Make it by chunks and not by players.

	p.PlayerList().Range(func(u uuid.UUID, player *Player) bool {
		if p.UUID() == u {
			return true
		}

		p.Conn().WritePacket(&play.SpawnPlayer{
			EntityID:    player.Eid.Load(),
			UUID:        player.UUID(),
			X:           player.Position().X(),
			Y:           player.Position().Y(),
			Z:           player.Position().Z(),
			YawAngle:    player.YawAngle(),
			PitchAngle:  player.PitchAngle(),
			CurrentItem: 1,
		})

		player.Conn().WritePacket(&play.SpawnPlayer{
			EntityID:    p.Eid.Load(),
			UUID:        p.UUID(),
			X:           p.Position().X(),
			Y:           p.Position().Y(),
			Z:           p.Position().Z(),
			YawAngle:    p.YawAngle(),
			PitchAngle:  p.PitchAngle(),
			CurrentItem: 1,
		})

		p.logger.Infof("Spawned %s to %s", p.Username(), player.Username())

		return true
	})
}

func (p *Player) LeaveGame() {
	p.PlayerList().Remove(p)

	p.PlayerList().BroadcastPacket(&play.DestroyEntities{
		EntityIDs: []int32{p.Eid.Load()},
	})
}

func (p *Player) HandleMove(x, y, z float64, yaw, pitch float32, onGround, moved, rotated bool) {
	prevFPX, prevFPY, prevFPZ := doubleToFixedPoint(p.Position().X()), doubleToFixedPoint(p.Position().Y()), doubleToFixedPoint(p.Position().Z())
	fpX, fpY, fpZ := doubleToFixedPoint(x), doubleToFixedPoint(y), doubleToFixedPoint(z)
	deltaFPX, deltaFPY, deltaFPZ := fpX-prevFPX, fpY-prevFPY, fpZ-prevFPZ

	fpFitsByte := func(fp int32) bool {
		return fp >= -128 && fp <= 127
	}

	if rotated {
		p.SetYaw(float32(math.Mod(float64(yaw)+360.0, 360.0)))
		p.SetPitch(pitch)
	}

	if moved {
		// TODO: Handle position change (loading/unloading chunks, etc.)

		p.SetPosition(mgl64.Vec3{x, y, z})
	}

	p.OnGround.Store(onGround)
	packetsToSend := make([]packet.Packet, 0)

	p.logger.Debugf("Player %s moved dX=%d dY=%d dZ=%d", p.Username(), deltaFPX, deltaFPY, deltaFPZ)

	if moved && fpFitsByte(deltaFPX) && fpFitsByte(deltaFPY) && fpFitsByte(deltaFPZ) {
		if rotated {
			packetsToSend = append(packetsToSend, &play.EntityRelativeMoveAndLook{
				EntityID: p.Eid.Load(),
				DeltaX:   byte(deltaFPX),
				DeltaY:   byte(deltaFPY),
				DeltaZ:   byte(deltaFPZ),
				Yaw:      yaw,
				Pitch:    pitch,
				OnGround: onGround,
			})
		} else {
			packetsToSend = append(packetsToSend, &play.EntityRelativeMove{
				EntityID: p.Eid.Load(),
				DeltaX:   byte(deltaFPX),
				DeltaY:   byte(deltaFPY),
				DeltaZ:   byte(deltaFPZ),
				OnGround: onGround,
			})
		}
	} else if rotated || moved {
		// TODO: Send Entity Teleport packet
	} else {
		// TODO: Send Entity Look packet
	}

	if rotated {
		// TODO: Send Entity Head Look packet
	}

	p.logger.Infof("Sent %d packets to other players", len(packetsToSend))
	p.playerList.Range(func(u uuid.UUID, player *Player) bool {
		if p.UUID() == u {
			return true
		}

		for _, p := range packetsToSend {
			player.Conn().WritePacket(p)
		}

		return true
	})
}

// TODO: Move this function to a more appropriate place.
func doubleToFixedPoint(val float64) int32 {
	return int32(val * 32.0)
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

func (p *Player) Position() mgl64.Vec3 {
	p.RLock()
	defer p.RUnlock()

	return p.position
}

func (p *Player) SetPosition(position mgl64.Vec3) {
	p.Lock()
	defer p.Unlock()

	p.position = position

	// TODO: Send position change packet to the client.
}

func (p *Player) Yaw() float32 {
	p.RLock()
	defer p.RUnlock()

	return p.yaw
}

func (p *Player) YawAngle() int8 {
	return int8(((math.Mod(float64(p.Yaw()), 360)) / 360) * 256)
}

func (p *Player) SetYaw(yaw float32) {
	p.Lock()
	defer p.Unlock()

	p.yaw = yaw

	// TODO: Send yaw change packet to the client.
}

func (p *Player) Pitch() float32 {
	p.RLock()
	defer p.RUnlock()

	return p.pitch
}

func (p *Player) PitchAngle() int8 {
	return int8(((math.Mod(float64(p.Pitch()), 360)) / 360) * 256)
}

func (p *Player) SetPitch(pitch float32) {
	p.Lock()
	defer p.Unlock()

	p.pitch = pitch
}

func (p *Player) Sneaking() bool {
	return p.sneaking.Load()
}

func (p *Player) SetSneaking(sneaking bool) {
	p.sneaking.Store(sneaking)

	// TODO: Send PlayEntityMetadata packet to other players.
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

func (p *Player) PlayerList() *PlayerList {
	return p.playerList
}
