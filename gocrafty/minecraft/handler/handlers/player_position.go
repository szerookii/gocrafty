package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type PlayerPositionHandler struct{}

func (h *PlayerPositionHandler) PacketID() int32 {
	return play.IDServerPlayerPosition
}

func (h *PlayerPositionHandler) Handle(p *player.Player, packet packet.Packet) {
	pkt, _ := packet.(*play.PlayerPosition)

	p.HandleMove(pkt.X, pkt.FeetY, pkt.Z, p.Yaw(), p.Pitch(), pkt.OnGround, true, false)
}
