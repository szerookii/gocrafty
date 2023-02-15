package status

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type PingRequest struct {
	PingTime int64
}

func (s *PingRequest) ID() int32 {
	return IDPing
}

func (s *PingRequest) State() int32 {
	return types.StateStatus
}

func (s *PingRequest) Marshal(w *protocol.Writer) {}

func (s *PingRequest) Unmarshal(r *protocol.Reader) {
	r.Long(&s.PingTime)
}
