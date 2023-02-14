package status

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type PingResponse struct {
	PingTime int64
}

func (s *PingResponse) ID() int32 {
	return packets.IDPing
}

func (s *PingResponse) State() int32 {
	return types.StateStatus
}

func (s *PingResponse) Marshal(w *protocol.Writer) {
	w.Long(s.PingTime)
}

func (s *PingResponse) Unmarshal(r *protocol.Reader) {}
