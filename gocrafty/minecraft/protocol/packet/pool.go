package packet

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/handshake"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/status"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

var (
	registeredHandshakePakets = map[int32]Packet{}
	registeredStatusPackets   = map[int32]Packet{}
	registeredLoginPackets    = map[int32]Packet{}
	registeredPlayPackets     = map[int32]Packet{}
)

func Register(id int32, pk Packet) {
	switch pk.State() {
	case types.StateHandshaking:
		registeredHandshakePakets[id] = pk
	case types.StateStatus:
		registeredStatusPackets[id] = pk
	case types.StateLogin:
		registeredLoginPackets[id] = pk
	case types.StatePlay:
		registeredPlayPackets[id] = pk
	}
}

type Pool struct {
	handshakePakets map[int32]Packet
	statusPackets   map[int32]Packet
	loginPackets    map[int32]Packet
	playPackets     map[int32]Packet
}

func NewPool() Pool {
	p := Pool{
		handshakePakets: map[int32]Packet{},
		statusPackets:   map[int32]Packet{},
		loginPackets:    map[int32]Packet{},
		playPackets:     map[int32]Packet{},
	}

	for id, pk := range registeredHandshakePakets {
		p.handshakePakets[id] = pk
	}

	for id, pk := range registeredStatusPackets {
		p.statusPackets[id] = pk
	}

	for id, pk := range registeredLoginPackets {
		p.loginPackets[id] = pk
	}

	for id, pk := range registeredPlayPackets {
		p.playPackets[id] = pk
	}

	return p
}

func (p Pool) Get(state int32, id int32) (Packet, bool) {
	switch state {
	case types.StateHandshaking:
		pk, ok := p.handshakePakets[id]
		return pk, ok
	case types.StateStatus:
		pk, ok := p.statusPackets[id]
		return pk, ok
	case types.StateLogin:
		pk, ok := p.loginPackets[id]
		return pk, ok
	case types.StatePlay:
		pk, ok := p.playPackets[id]
		return pk, ok
	}

	return nil, false
}

func init() {
	Register(packets.IDHandshake, &handshake.Handshake{})

	Register(packets.IDStatusRequest, &status.StatusRequest{})
}
