package handlers

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type EntityActionHandler struct{}

func (h *EntityActionHandler) PacketID() int32 {
	return play.IDServerEntityAction
}

func (h *EntityActionHandler) Handle(p *player.Player, packet packet.Packet) {
	pkt, _ := packet.(*play.EntityAction)

	switch pkt.ActionID {
	case play.ActionStartSneaking:
		p.SetSneaking(true)
	case play.ActionStopSneaking:
		p.SetSneaking(false)
	case play.ActionLeaveBed:
		// TODO: Implement
	case play.ActionStartSprinting:
		p.Sprinting.Store(true)
	case play.ActionStopSprinting:
		p.Sprinting.Store(false)
	case play.ActionJumpWithHorse:
	// TODO: Implement
	case play.ActionOpenHorseInventory:
		// TODO: Implement

	}
}
