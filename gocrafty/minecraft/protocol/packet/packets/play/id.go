package play

const (
	// Clientbound
	IDClientJoinGame              = 0x01
	IDClientChatMessage           = 0x02
	IDClientTimeUpdate            = 0x03
	IDClientPlayerPositionAndLook = 0x08
	IDClientAnimation             = 0x0B
	IDClientSpawnPlayer           = 0x0C
	IDClientDestroyEntities       = 0x13
	IDClientChunkData             = 0x21
	IDClientPlayerListItem        = 0x38
	IDClientDisconnect            = 0x40

	// Serverbound
	IDServerKeepAlive                 = 0x00
	IDServerChatMessage               = 0x01
	IDServerUseEntity                 = 0x02
	IDServerPlayer                    = 0x03
	IDServerPlayerPosition            = 0x04
	IDServerPlayerLook                = 0x05
	IDServerPlayerPositionAndLook     = 0x06
	IDServerAnimation                 = 0x0A
	IDServerEntityAction              = 0x0B
	IDServerEntityRelativeMove        = 0x15
	IDServerEntityLook                = 0x16
	IDServerEntityRelativeMoveAndLook = 0x17
	IDServerEntityTeleport            = 0x18
	IDServerEntityHeadLook            = 0x19
)
