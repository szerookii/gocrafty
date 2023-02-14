package handler

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/session"
	"github.com/szerookii/gocrafty/gocrafty/session/handler/handshake"
)

type HandlerData struct {
	Packet  packet.Packet
	Session *session.Session
}

type Handler interface {
	PacketID() int32
	Handshake() bool
	Handle(data HandlerData) error
}

var handlers = make(map[int32]Handler)

func RegisterHandler(handler Handler) {
	handlers[handler.PacketID()] = handler
}

func GetHandler(packetID int32) (Handler, bool) {
	handler, ok := handlers[packetID]
	return handler, ok
}

func init() {
	RegisterHandler(&handshake.HandshakeHandler{})
}
