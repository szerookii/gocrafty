package play

import (
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

// Définition des actions
const (
	PlayerListItemActionAddPlayer         = 0
	PlayerListItemActionUpdateGamemode    = 1
	PlayerListItemActionUpdateLatency     = 2
	PlayerListItemActionUpdateDisplayName = 3
	PlayerListItemActionRemovePlayer      = 4
)

type PlayerListItemDataInterface interface {
	ID() int32
	Marshal(w *protocol.Writer)
}

type PlayerListItemData struct {
	UUID uuid.UUID
}

type PlayerListItemDataAddPlayer struct {
	PlayerListItemData
	Name       string
	Properties []struct {
		Name      string
		Value     string
		Signature string
	}
	Gamemode       int32
	Latency        int32
	HasDisplayName bool
	DisplayName    string
}

type PlayerListItemDataUpdateGamemode struct {
	PlayerListItemData
	Gamemode int32
}

type PlayerListItemDataUpdateLatency struct {
	PlayerListItemData
	Latency int32
}

type PlayerListItemDataUpdateDisplayName struct {
	PlayerListItemData
	HasDisplayName bool
	DisplayName    string
}

type PlayerListItemDataRemovePlayer struct {
	PlayerListItemData
}

func (p *PlayerListItemDataAddPlayer) ID() int32 {
	return PlayerListItemActionAddPlayer
}

func (p *PlayerListItemDataUpdateGamemode) ID() int32 {
	return PlayerListItemActionUpdateGamemode
}

func (p *PlayerListItemDataUpdateLatency) ID() int32 {
	return PlayerListItemActionUpdateLatency
}

func (p *PlayerListItemDataUpdateDisplayName) ID() int32 {
	return PlayerListItemActionUpdateDisplayName
}

func (p *PlayerListItemDataRemovePlayer) ID() int32 {
	return PlayerListItemActionRemovePlayer
}

func (p *PlayerListItemDataAddPlayer) Marshal(w *protocol.Writer) {
	w.UUID(p.UUID)
	w.String(p.Name)
	w.VarInt(int32(len(p.Properties)))
	for _, property := range p.Properties {
		w.String(property.Name)
		w.String(property.Value)
		w.String(property.Signature)
	}
	w.VarInt(p.Gamemode)
	w.VarInt(p.Latency)
	w.Bool(p.HasDisplayName)
	if p.HasDisplayName {
		w.String(p.DisplayName)
	}
}

func (p *PlayerListItemDataUpdateGamemode) Marshal(w *protocol.Writer) {
	w.UUID(p.UUID)
	w.VarInt(p.Gamemode)
}

func (p *PlayerListItemDataUpdateLatency) Marshal(w *protocol.Writer) {
	w.UUID(p.UUID)
	w.VarInt(p.Latency)
}

func (p *PlayerListItemDataUpdateDisplayName) Marshal(w *protocol.Writer) {
	w.UUID(p.UUID)
	w.Bool(p.HasDisplayName)
	if p.HasDisplayName {
		w.String(p.DisplayName)
	}
}

func (p *PlayerListItemDataRemovePlayer) Marshal(w *protocol.Writer) {
	w.UUID(p.UUID)
}

type PlayerListItem struct {
	Action  int32
	Players []PlayerListItemDataInterface
}

func (p *PlayerListItem) ID() int32 {
	return IDClientPlayerListItem
}

func (p *PlayerListItem) State() int32 {
	return types.StatePlay
}

func (p *PlayerListItem) Marshal(w *protocol.Writer) {
	w.VarInt(p.Action)
	w.VarInt(int32(len(p.Players)))
	for _, player := range p.Players {
		player.Marshal(w)
	}
}

func (p *PlayerListItem) Unmarshal(_ *protocol.Reader) {}
