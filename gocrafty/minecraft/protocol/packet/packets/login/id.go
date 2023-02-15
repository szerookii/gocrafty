package login

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
