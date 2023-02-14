package session

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft"
	"github.com/szerookii/gocrafty/gocrafty/session/handler"
)

type Session struct {
	conn     *minecraft.Conn
	handlers map[int32]handler.Handler
}
