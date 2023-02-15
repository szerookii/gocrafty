package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type KeepAliveHandler struct{}

func (h *KeepAliveHandler) PacketID() int32 {
	return play.IDKeepAlive
}

func (h *KeepAliveHandler) Handle(p *player.Player, packet packet.Packet) {
	keepAlive, _ := packet.(*play.KeepAlive)

	p.Logger().Debugf("Received keep alive packet with ID %d", keepAlive.KeepAliveId)
}
