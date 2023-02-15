package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type Handler interface {
	PacketID() int32
	Handle(p *player.Player, packet packet.Packet)
}
