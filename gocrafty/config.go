package gocrafty

import "github.com/szerookii/gocrafty/gocrafty/logger"

type ServerConfig struct {
	// Logger is the logger used by the server.
	Logger   logger.Logger `json:"-"`
	LogLevel string        `json:"log_level"`

	// BoundAddress is the address the server will bind to.
	BoundAddress string `json:"bound_address"`

	// MaxPlayers is the maximum amount of players that can be connected to the server at the same time.
	MaxPlayers int `json:"max_players"`
	// OnlineMode is whether the server will use online mode or not.
	OnlineMode bool `json:"online_mode"`
	// ServerName is the name of the server.
	ServerName string `json:"server_name"`
}

func DefaultConfig() ServerConfig {
	return ServerConfig{
		Logger:       logger.Default(),
		LogLevel:     "info",
		BoundAddress: ":25565",

		MaxPlayers: 20,
		OnlineMode: true,
		ServerName: "A Gocrafty Server",
	}
}
