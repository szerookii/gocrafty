package handlers

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type AnimationHandler struct{}

func (h *AnimationHandler) PacketID() int32 {
	return play.IDServerAnimation
}

func (h *AnimationHandler) Handle(p *player.Player, packet packet.Packet) {
	p.PlayerList().Range(func(uuid uuid.UUID, op *player.Player) bool {
		if op.Eid.Load() == p.Eid.Load() {
			return true
		}

		op.Conn().WritePacket(&play.ClientAnimation{
			EntityID:    p.Eid.Load(),
			AnimationID: 0, // Swing arm
		})

		return true
	})
}
