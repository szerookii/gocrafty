package login

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type LoginStart struct {
	Username string
}

func (d *LoginStart) ID() int32 {
	return packets.IDLoginStart
}

func (s *LoginStart) State() int32 {
	return types.StateLogin
}

func (d *LoginStart) Unmarshal(r *protocol.Reader) {
	r.String(&d.Username)
}

func (d *LoginStart) Marshal(w *protocol.Writer) {}
