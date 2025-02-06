package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type PlayerPositionAndLookHandler struct{}

func (h *PlayerPositionAndLookHandler) PacketID() int32 {
	return play.IDClientPlayerPositionAndLook
}

func (h *PlayerPositionAndLookHandler) Handle(p *player.Player, packet packet.Packet) {
	pkt, _ := packet.(*play.ServerPlayerPositionAndLook)
	
	p.HandleMove(pkt.X, pkt.FeetY, pkt.Z, pkt.Yaw, pkt.Pitch, pkt.OnGround, true, true)
}
