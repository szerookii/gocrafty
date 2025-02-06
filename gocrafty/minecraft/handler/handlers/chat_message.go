package handlers

import (
	"fmt"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
	"github.com/szerookii/gocrafty/gocrafty/player"
)

type ChatMessageHandler struct{}

func (h *ChatMessageHandler) PacketID() int32 {
	return play.IDServerChatMessage
}

func (h *ChatMessageHandler) Handle(p *player.Player, packet packet.Packet) {
	pkt, _ := packet.(*play.ServerChatMessage)

	p.Logger().Infof("%s > %s", p.Username(), pkt.Message)

	// TODO: Parse the message and handle commands
	p.PlayerList().BroadcastPacket(&play.ClientChatMessage{
		Data: types.Chat{
			Text: fmt.Sprintf("%s > %s", p.Username(), pkt.Message),
		},
		Position: play.ClientChatMessagePositionChat,
	})
}
