package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type EntityRelativeMove struct {
	EntityID int32
	DeltaX   byte
	DeltaY   byte
	DeltaZ   byte
	OnGround bool
}

func (p *EntityRelativeMove) ID() int32 {
	return IDServerEntityRelativeMove
}

func (p *EntityRelativeMove) State() int32 {
	return types.StatePlay
}

func (p *EntityRelativeMove) Marshal(w *protocol.Writer) {
	w.VarInt(p.EntityID)
	w.WriteByte(p.DeltaX)
	w.WriteByte(p.DeltaY)
	w.WriteByte(p.DeltaZ)
	w.Bool(p.OnGround)
}

func (p *EntityRelativeMove) Unmarshal(r *protocol.Reader) {}
