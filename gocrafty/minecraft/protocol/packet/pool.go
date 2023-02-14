package packet

var registeredPackets = map[int32]func(bool) Packet{}

func Register(id int32, pk func(bool) Packet) {
	registeredPackets[id] = pk
}

type Pool struct {
	handshakePakets  map[int32]Packet
	connectedPackets map[int32]Packet
}

func NewPool() Pool {
	p := Pool{
		handshakePakets:  map[int32]Packet{},
		connectedPackets: map[int32]Packet{},
	}

	for id, pk := range registeredPackets {
		if pk(false) != nil {
			p.handshakePakets[id] = pk(false)
		}

		if pk(true) != nil {
			p.connectedPackets[id] = pk(true)
		}
	}

	return p
}

func (p Pool) Get(handshaked bool, id int32) (Packet, bool) {
	if handshaked {
		pk, ok := p.connectedPackets[id]
		return pk, ok
	}

	pk, ok := p.handshakePakets[id]
	return pk, ok
}

func init() {
	pkts := map[int32]func(bool) Packet{
		// Handshake
		IDHandshake: func(connected bool) Packet {
			if connected {
				return nil
			}

			return &Handshake{}
		},
	}

	for id, pk := range pkts {
		Register(id, pk)
	}
}
