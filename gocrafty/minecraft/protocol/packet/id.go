package packet

// Hanshaking packets.
const (
	// Handshake is the ID of the Handshake packet.
	IDHandshake = 0x00
	// LegacyServerListPing is the ID of the LegacyServerListPing packet.
	IDLegacyServerListPing = 0xfe
)

// Play packets.
const (
	// KeepAlive is the ID of the KeepAlive packet.
	IDKeepAlive = 0x00
	// JoinGame is the ID of the JoinGame packet.
	IDJoinGame = 0x01
	// ChatMessage is the ID of the ChatMessage packet.
	IDChatMessage = 0x02
	// TimeUpdate is the ID of the TimeUpdate packet.
	IDTimeUpdate = 0x03
)
