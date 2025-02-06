package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type KeepAlive struct {
	KeepAliveId int32
}

func (k *KeepAlive) ID() int32 {
	return IDServerKeepAlive
}

func (k *KeepAlive) State() int32 {
	return types.StatePlay
}

func (k *KeepAlive) Marshal(w *protocol.Writer) {
	w.VarInt(k.KeepAliveId)
}

func (k *KeepAlive) Unmarshal(r *protocol.Reader) {
	r.VarInt(&k.KeepAliveId)
}
