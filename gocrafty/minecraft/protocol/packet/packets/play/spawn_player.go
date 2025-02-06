package play

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type SpawnPlayer struct {
	EntityID    int32
	UUID        uuid.UUID
	X           float64
	Y           float64
	Z           float64
	YawAngle    int8
	PitchAngle  int8
	CurrentItem int16
	// TODO: Implement metadata
}

func (s *SpawnPlayer) ID() int32 {
	return IDClientSpawnPlayer
}

func (s *SpawnPlayer) State() int32 {
	return types.StatePlay
}

func (s *SpawnPlayer) Marshal(w *protocol.Writer) {
	w.VarInt(s.EntityID)
	w.UUID(s.UUID)
	w.FixedPoint(s.X)
	w.FixedPoint(s.Y)
	w.FixedPoint(s.Z)
	w.WriteByte(byte(s.YawAngle))
	w.WriteByte(byte(s.PitchAngle))
	w.Short(s.CurrentItem)
	w.WriteBytes([]byte{0, 0x08, 0x7F}) //TODO: Implement metadata
}

func (s *SpawnPlayer) Unmarshal(_ *protocol.Reader) {}
