package handler

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/handler/handlers"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
)

func GetHandler(id int32) (handlers.Handler, bool) {
	handlersMap := map[int32]handlers.Handler{
		play.IDKeepAlive: &handlers.KeepAliveHandler{},
	}

	if h, ok := handlersMap[id]; ok {
		return h, true
	}

	return nil, false
}
