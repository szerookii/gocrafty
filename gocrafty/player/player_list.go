package player

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"sync"
)

type PlayerList struct {
	mu      *sync.RWMutex
	players map[uuid.UUID]*Player
}

func NewPlayerList() *PlayerList {
	return &PlayerList{
		mu:      &sync.RWMutex{},
		players: make(map[uuid.UUID]*Player),
	}
}

func (pl *PlayerList) BroadcastPacket(p packet.Packet) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	for _, player := range pl.players {
		player.Conn().WritePacket(p)
	}
}

func (pl *PlayerList) Add(p *Player) {
	pl.mu.Lock()
	pl.players[p.UUID()] = p
	pl.mu.Unlock()

	p.Conn().WritePacket(pl.ListPacket())

	pl.Range(func(uuid uuid.UUID, player *Player) bool {
		if p.UUID() == uuid {
			return true
		}

		player.Conn().WritePacket(&play.PlayerListItem{
			Action: play.PlayerListItemActionAddPlayer,
			Players: []play.PlayerListItemDataInterface{
				&play.PlayerListItemDataAddPlayer{
					PlayerListItemData: play.PlayerListItemData{
						UUID: p.UUID(),
					},
					Name:           p.Username(),
					Properties:     nil,
					Gamemode:       int32(p.Gamemode()),
					Latency:        0,
					HasDisplayName: false,
				},
			},
		})

		return true
	})
}

func (pl *PlayerList) Remove(p *Player) {
	pl.mu.Lock()
	delete(pl.players, p.UUID())
	pl.mu.Unlock()

	pl.BroadcastPacket(&play.PlayerListItem{
		Action: play.PlayerListItemActionRemovePlayer,
		Players: []play.PlayerListItemDataInterface{
			&play.PlayerListItemDataRemovePlayer{
				PlayerListItemData: play.PlayerListItemData{
					UUID: p.UUID(),
				},
			},
		},
	})
}

func (pl *PlayerList) ListPacket() *play.PlayerListItem {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	var players []play.PlayerListItemDataInterface
	for _, p := range pl.players {
		players = append(players, &play.PlayerListItemDataAddPlayer{
			PlayerListItemData: play.PlayerListItemData{
				UUID: p.UUID(),
			},
			Name:           p.Username(),
			Properties:     nil,
			Gamemode:       int32(p.Gamemode()),
			Latency:        0,
			HasDisplayName: false,
		})
	}

	return &play.PlayerListItem{
		Action:  play.PlayerListItemActionAddPlayer,
		Players: players,
	}
}

func (pl *PlayerList) Get(uuid uuid.UUID) (*Player, bool) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	p, ok := pl.players[uuid]
	return p, ok
}

func (pl *PlayerList) GetByName(name string) (*Player, bool) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	for _, p := range pl.players {
		if p.Username() == name {
			return p, true
		}
	}

	return nil, false
}

func (pl *PlayerList) All() map[uuid.UUID]*Player {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	return pl.players
}

func (pl *PlayerList) Len() int {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	return len(pl.players)
}

func (pl *PlayerList) Range(f func(uuid.UUID, *Player) bool) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	for uuid, p := range pl.players {
		if !f(uuid, p) {
			break
		}
	}
}

func (pl *PlayerList) Close() {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	for _, p := range pl.players {
		p.Disconnect("Server shutting down")
	}

	pl.players = make(map[uuid.UUID]*Player)
}
