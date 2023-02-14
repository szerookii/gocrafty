package packet

import "github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"

type Packet interface {
	// ID returns the ID of the packet.
	ID() int32
	// Marshal marshals the packet into a byte slice.
	Marshal(w *protocol.Writer)
	// Unmarshal unmarshals the packet from a byte slice.
	Unmarshal(r *protocol.Reader)
}
