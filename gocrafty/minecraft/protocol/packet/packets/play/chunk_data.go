package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type ChunkData struct {
	ChunkX, ChunkZ     int32
	GroundUpContinuous bool
	PrimaryBitMask     uint16
	Data               []byte
}

func (c *ChunkData) ID() int32 {
	return IDClientChunkData
}

func (c *ChunkData) State() int32 {
	return types.StatePlay
}

func (c *ChunkData) Marshal(w *protocol.Writer) {
	w.Int(c.ChunkX)
	w.Int(c.ChunkZ)
	w.Bool(c.GroundUpContinuous)
	w.UShort(c.PrimaryBitMask)
	w.VarInt(int32(len(c.Data)))
	w.WriteBytes(c.Data)
}

func (c *ChunkData) Unmarshal(r *protocol.Reader) {
	r.Int(&c.ChunkX)
	r.Int(&c.ChunkZ)
	r.Bool(&c.GroundUpContinuous)
	r.UShort(&c.PrimaryBitMask)
	var length int32
	r.VarInt(&length)
	c.Data = r.ReadBytes(int(length))
}
