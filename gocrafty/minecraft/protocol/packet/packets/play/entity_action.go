package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

const (
	ActionStartSneaking = iota
	ActionStopSneaking
	ActionLeaveBed
	ActionStartSprinting
	ActionStopSprinting
	ActionJumpWithHorse
	ActionOpenHorseInventory
)

type EntityAction struct {
	EntityID int32
	ActionID byte
}

func (a *EntityAction) ID() int32 {
	return IDServerEntityAction
}

func (a *EntityAction) State() int32 {
	return types.StatePlay
}

func (a *EntityAction) Marshal(w *protocol.Writer) {
	w.VarInt(a.EntityID)
	w.WriteByte(a.ActionID)
}

func (a *EntityAction) Unmarshal(r *protocol.Reader) {
	r.VarInt(&a.EntityID)
	r.Byte(&a.ActionID)
}
