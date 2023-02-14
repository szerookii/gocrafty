package handshake

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/session/handler"
)

type HandshakeHandler struct{}

func (h *HandshakeHandler) PacketID() int32 {
	return packets.IDHandshake
}

func (h *HandshakeHandler) Handshake() bool {
	return true
}

func (h *HandshakeHandler) Handle(data handler.HandlerData) error {
	return nil
}
