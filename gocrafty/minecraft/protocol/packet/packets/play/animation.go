package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type ServerAnimation struct{}

func (a *ServerAnimation) ID() int32 {
	return IDServerAnimation
}

func (a *ServerAnimation) State() int32 {
	return types.StatePlay
}

func (a *ServerAnimation) Marshal(w *protocol.Writer) {}

func (a *ServerAnimation) Unmarshal(r *protocol.Reader) {}

type ClientAnimation struct {
	EntityID    int32
	AnimationID byte
}

func (a *ClientAnimation) ID() int32 {
	return IDClientAnimation
}

func (a *ClientAnimation) State() int32 {
	return types.StatePlay
}

func (a *ClientAnimation) Marshal(w *protocol.Writer) {
	w.VarInt(a.EntityID)
	w.WriteByte(a.AnimationID)
}

func (a *ClientAnimation) Unmarshal(r *protocol.Reader) {
	r.VarInt(&a.EntityID)
	r.Byte(&a.AnimationID)
}
