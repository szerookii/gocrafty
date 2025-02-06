package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type PlayerLook struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *PlayerLook) ID() int32 {
	return IDServerPlayerLook
}

func (p *PlayerLook) State() int32 {
	return types.StatePlay
}

func (p *PlayerLook) Marshal(*protocol.Writer) {}

func (p *PlayerLook) Unmarshal(r *protocol.Reader) {
	r.Float(&p.Yaw)
	r.Float(&p.Pitch)
	r.Bool(&p.OnGround)
}
