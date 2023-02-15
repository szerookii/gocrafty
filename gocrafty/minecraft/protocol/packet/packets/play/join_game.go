package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type JoinGame struct {
	EID              int32
	Gamemode         uint8
	Dimension        byte
	Difficulty       uint8
	MaxPlayers       uint8
	LevelType        string
	ReducedDebugInfo bool
}

func (j *JoinGame) ID() int32 {
	return IDJoinGame
}

func (j *JoinGame) State() int32 {
	return types.StatePlay
}

func (j *JoinGame) Marshal(w *protocol.Writer) {
	w.Int(j.EID)
	w.WriteByte(j.Gamemode)
	w.WriteByte(j.Dimension)
	w.WriteByte(j.Difficulty)
	w.WriteByte(j.MaxPlayers)
	w.String(j.LevelType)
	w.Bool(j.ReducedDebugInfo)
}

func (j *JoinGame) Unmarshal(r *protocol.Reader) {
	r.Int(&j.EID)
	r.Byte(&j.Gamemode)
	r.Byte(&j.Dimension)
	r.Byte(&j.Difficulty)
	r.Byte(&j.MaxPlayers)
	r.String(&j.LevelType)
	r.Bool(&j.ReducedDebugInfo)
}
