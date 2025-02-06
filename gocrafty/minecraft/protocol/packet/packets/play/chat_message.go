package play

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

const (
	ClientChatMessagePositionChat = iota
	ClientChatMessagePositionSystemMessage
	ClientChatMessageAboveHotbar
)

type ServerChatMessage struct {
	Message string
}

func (s *ServerChatMessage) ID() int32 {
	return IDServerChatMessage
}

func (s *ServerChatMessage) State() int32 {
	return types.StatePlay
}

func (s *ServerChatMessage) Marshal(_ *protocol.Writer) {}

func (s *ServerChatMessage) Unmarshal(r *protocol.Reader) {
	r.String(&s.Message)
}

type ClientChatMessage struct {
	Data     types.Chat
	Position byte
}

func (c *ClientChatMessage) ID() int32 {
	return IDClientChatMessage
}

func (c *ClientChatMessage) State() int32 {
	return types.StatePlay
}

func (c *ClientChatMessage) Marshal(w *protocol.Writer) {
	w.Chat(c.Data)
	w.WriteByte(c.Position)
}

func (c *ClientChatMessage) Unmarshal(_ *protocol.Reader) {}
