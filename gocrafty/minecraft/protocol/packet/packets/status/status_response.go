package status

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
)

type StatusResponseDataVersion struct {
	// Name is the name of the version.
	Name string `json:"name"`
	// Protocol is the protocol version of the server.
	Protocol int32 `json:"protocol"`
}

type StatusResponseDataPlayers struct {
	// Max is the maximum amount of players that can be connected to the server at the same time.
	Max int32 `json:"max"`
	// Online is the amount of players that are currently connected to the server.
	Online int32 `json:"online"`
}

type StatusResponseDataDescription struct {
	Text string `json:"text"`
}

type StatusResponseData struct {
	// Version is the version of the server.
	Version *StatusResponseDataVersion `json:"version"`
	// Players is the amount of players that are currently connected to the server.
	Players *StatusResponseDataPlayers `json:"players"`
	// Description is the description of the server.
	Description *StatusResponseDataDescription `json:"description"`
	// Favicon is the base64 encoded favicon of the server.
	Favicon string `json:"favicon,omitempty"`
}

// StatusResponse is sent by the server to the client in response to a StatusRequest packet.
type StatusResponse struct {
	JSONResponse *StatusResponseData
}

func (s *StatusResponse) ID() int32 {
	return packets.IDStatusResponse
}

func (s *StatusResponse) State() int32 {
	return types.StateLogin
}

func (s *StatusResponse) Marshal(w *protocol.Writer) {
	w.JSON(s.JSONResponse)
}

func (s *StatusResponse) Unmarshal(r *protocol.Reader) {}
