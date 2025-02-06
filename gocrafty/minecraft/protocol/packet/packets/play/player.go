package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type Player struct {
	OnGround bool
}

func (p *Player) ID() int32 {
	return IDServerPlayer
}

func (p *Player) State() int32 {
	return types.StatePlay
}

func (p *Player) Marshal(w *protocol.Writer) {
	w.Bool(p.OnGround)
}

func (p *Player) Unmarshal(r *protocol.Reader) {
	r.Bool(&p.OnGround)
}
