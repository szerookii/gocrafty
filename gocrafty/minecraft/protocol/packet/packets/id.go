package packets

// Hanshaking packets.
const (
	// Handshake is the ID of the Handshake packet.
	IDHandshake = 0x00
	// LegacyServerListPing is the ID of the LegacyServerListPing packet.
	IDLegacyServerListPing = 0xfe
)

// Status packets.
const (
	// StatusRequest is the ID of the StatusRequest packet.
	IDStatusRequest = 0x00
	// StatusResponse is the ID of the StatusResponse packet.
	IDStatusResponse = 0x00
	// Ping is the ID of the Ping packet.
	IDPing = 0x01
)

// Login packets.
const (
	// Disconnect is the ID of the Disconnect packet.
	IDDisconnect = 0x00
	// EncryptionRequest is the ID of the EncryptionRequest packet.
	IDEncryptionRequest = 0x01
	// LoginSuccess is the ID of the LoginSuccess packet.
	IDLoginSuccess = 0x02
	// SetCompression is the ID of the SetCompression packet.
	IDSetCompression = 0x03
	// LoginPluginRequest is the ID of the LoginPluginRequest packet.
	IDLoginPluginRequest = 0x04

	// LoginStart is the ID of the LoginStart packet.
	IDLoginStart = 0x00
	// EncryptionResponse is the ID of the EncryptionResponse packet.
	IDEncryptionResponse = 0x01
	// LoginPluginResponse is the ID of the LoginPluginResponse packet.
	IDLoginPluginResponse = 0x02
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
