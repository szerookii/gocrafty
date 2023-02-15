package packet

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/handshake"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/login"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/status"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type Pool struct {
	handshakePackets, statusPackets, loginPackets, playPackets map[int32]Packet
}

func NewPool() Pool {
	return Pool{
		handshakePackets: make(map[int32]Packet),
		statusPackets:    make(map[int32]Packet),
		loginPackets:     make(map[int32]Packet),
		playPackets:      make(map[int32]Packet),
	}
}

func (p Pool) Get(state, id int32) (Packet, bool) {
	switch state {
	case types.StateHandshaking:
		if pk, ok := p.handshakePackets[id]; ok {
			return pk, true
		}
	case types.StateStatus:
		if pk, ok := p.statusPackets[id]; ok {
			return pk, true
		}
	case types.StateLogin:
		if pk, ok := p.loginPackets[id]; ok {
			return pk, true
		}
	case types.StatePlay:
		if pk, ok := p.playPackets[id]; ok {
			return pk, true
		}
	}
	return nil, false
}

func Register(id int32, pk Packet) {
	switch pk.State() {
	case types.StateHandshaking:
		pool.handshakePackets[id] = pk
	case types.StateStatus:
		pool.statusPackets[id] = pk
	case types.StateLogin:
		pool.loginPackets[id] = pk
	case types.StatePlay:
		pool.playPackets[id] = pk
	}
}

func init() {
	Register(packets.IDHandshake, &handshake.Handshake{})
	Register(packets.IDStatusRequest, &status.StatusRequest{})
	Register(packets.IDPing, &status.PingRequest{})
	Register(packets.IDLoginStart, &login.LoginStart{})
}
