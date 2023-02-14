package handshake

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type Disconnect struct {
	Reason *types.Chat
}

func (d *Disconnect) Unmarshal(r *protocol.Reader) {
	panic("implement me")
}

func (d *Disconnect) ID() int32 {
	return packets.IDDisconnect
}

func (s *Disconnect) State() int32 {
	return 0 // not used
}

func (d *Disconnect) Marshal(w *protocol.Writer) {
	w.Chat(d.Reason)
}
