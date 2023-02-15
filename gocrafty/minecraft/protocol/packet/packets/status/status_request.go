package status

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type StatusRequest struct{}

func (s *StatusRequest) ID() int32 {
	return IDStatusRequest
}

func (s *StatusRequest) State() int32 {
	return types.StateStatus
}

func (s *StatusRequest) Marshal(w *protocol.Writer) {}

func (s *StatusRequest) Unmarshal(r *protocol.Reader) {}
