package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type EntityRelativeMoveAndLook struct {
	EntityID int32
	DeltaX   byte
	DeltaY   byte
	DeltaZ   byte
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *EntityRelativeMoveAndLook) ID() int32 {
	return IDServerEntityRelativeMoveAndLook
}

func (p *EntityRelativeMoveAndLook) State() int32 {
	return types.StatePlay
}

func (p *EntityRelativeMoveAndLook) Marshal(w *protocol.Writer) {
	w.VarInt(p.EntityID)
	w.WriteByte(p.DeltaX)
	w.WriteByte(p.DeltaY)
	w.WriteByte(p.DeltaZ)
	w.Angle(p.Yaw)
	w.Angle(p.Pitch)
	w.Bool(p.OnGround)
}

func (p *EntityRelativeMoveAndLook) Unmarshal(r *protocol.Reader) {}
