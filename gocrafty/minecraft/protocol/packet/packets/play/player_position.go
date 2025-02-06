package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type PlayerPosition struct {
	X        float64
	FeetY    float64
	Z        float64
	OnGround bool
}

func (p *PlayerPosition) ID() int32 {
	return IDServerPlayerPosition
}

func (p *PlayerPosition) State() int32 {
	return types.StatePlay
}

func (p *PlayerPosition) Marshal(_ *protocol.Writer) {}

func (p *PlayerPosition) Unmarshal(r *protocol.Reader) {
	r.Double(&p.X)
	r.Double(&p.FeetY)
	r.Double(&p.Z)
	r.Bool(&p.OnGround)
}
