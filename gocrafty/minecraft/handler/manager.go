package handler

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/handler/handlers"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
)

func GetHandler(id int32) (handlers.Handler, bool) {
	handlersMap := map[int32]handlers.Handler{
		play.IDServerKeepAlive:      &handlers.KeepAliveHandler{},
		play.IDServerPlayer:         &handlers.PlayerHandler{},
		play.IDServerAnimation:      &handlers.AnimationHandler{},
		play.IDServerEntityAction:   &handlers.EntityActionHandler{},
		play.IDServerChatMessage:    &handlers.ChatMessageHandler{},
		play.IDServerPlayerPosition: &handlers.PlayerPositionHandler{},
		// TODO: Player Look
		play.IDServerPlayerPositionAndLook: &handlers.PlayerPositionAndLookHandler{},
	}

	if h, ok := handlersMap[id]; ok {
		return h, true
	}

	return nil, false
}
