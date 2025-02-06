package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type DestroyEntities struct {
	EntityIDs []int32
}

func (p *DestroyEntities) ID() int32 {
	return IDClientDestroyEntities
}

func (p *DestroyEntities) State() int32 {
	return types.StatePlay
}

func (p *DestroyEntities) Marshal(w *protocol.Writer) {
	w.VarInt(int32(len(p.EntityIDs)))
	for _, id := range p.EntityIDs {
		w.VarInt(id)
	}
}

func (p *DestroyEntities) Unmarshal(r *protocol.Reader) {}
