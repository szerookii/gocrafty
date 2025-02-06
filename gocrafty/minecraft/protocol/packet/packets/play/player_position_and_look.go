package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type ServerPlayerPositionAndLook struct {
	X        float64
	FeetY    float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *ServerPlayerPositionAndLook) ID() int32 {
	return IDServerPlayerPositionAndLook
}

func (p *ServerPlayerPositionAndLook) State() int32 {
	return types.StatePlay
}

func (p *ServerPlayerPositionAndLook) Marshal(w *protocol.Writer) {}

func (p *ServerPlayerPositionAndLook) Unmarshal(r *protocol.Reader) {
	r.Double(&p.X)
	r.Double(&p.FeetY)
	r.Double(&p.Z)
	r.Float(&p.Yaw)
	r.Float(&p.Pitch)
	r.Bool(&p.OnGround)
}

type ClientPlayerPositionAndLook struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
	Flags byte
}

func (p *ClientPlayerPositionAndLook) ID() int32 {
	return IDClientPlayerPositionAndLook
}

func (p *ClientPlayerPositionAndLook) State() int32 {
	return types.StatePlay
}

func (p *ClientPlayerPositionAndLook) Marshal(w *protocol.Writer) {
	w.Double(p.X)
	w.Double(p.Y)
	w.Double(p.Z)
	w.Float(p.Yaw)
	w.Float(p.Pitch)
	w.WriteByte(p.Flags)
}

func (p *ClientPlayerPositionAndLook) Unmarshal(r *protocol.Reader) {
	r.Double(&p.X)
	r.Double(&p.Y)
	r.Double(&p.Z)
	r.Float(&p.Yaw)
	r.Float(&p.Pitch)
	r.Byte(&p.Flags)
}
