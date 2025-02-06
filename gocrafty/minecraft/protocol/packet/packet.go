package packet

import "github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"

type Packet interface {
	// ID returns the ID of the packet.
	ID() int32
	State() int32
	// Marshal marshals the packet into a byte slice.
	Marshal(w *protocol.Writer)
	// Unmarshal unmarshals the packet from a byte slice.
	Unmarshal(r *protocol.Reader)
}

type RawPacket struct {
	IDField int32
}

func (p *RawPacket) ID() int32 {
	return p.IDField
}

func (p *RawPacket) State() int32 {
	return -1
}

func (p *RawPacket) Marshal(_ *protocol.Writer) {}

func (p *RawPacket) Unmarshal(_ *protocol.Reader) {}
