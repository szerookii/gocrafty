package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type PlayerHandler struct{}

func (h *PlayerHandler) PacketID() int32 {
	return play.IDServerPlayer
}

func (h *PlayerHandler) Handle(p *player.Player, packet packet.Packet) {
	pkt, _ := packet.(*play.Player)

	p.HandleMove(p.Position().X(), p.Position().Y(), p.Position().Z(), p.Yaw(), p.Pitch(), pkt.OnGround, false, false)
}
