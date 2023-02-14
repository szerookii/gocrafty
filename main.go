package main

import (
	"encoding/json"
	"fmt"
	"github.com/szerookii/gocrafty/gocrafty"
	"github.com/szerookii/gocrafty/gocrafty/logger"
	"github.com/szerookii/gocrafty/gocrafty/minecraft"
	"os"
)

func main() {
	log := logger.Default()

	log.Infof("Starting Gocrafty for Minecraft v%s", minecraft.Version)

	conf, err := readConfig()
	if err != nil {
		log.Fatal("read config: %v", err)
	}

	srv := gocrafty.NewServer(conf)

	err = srv.Listen()
	if err != nil {
		log.Fatal("listen: %v", err)
	}

	for {
	}
}

func readConfig() (*gocrafty.ServerConfig, error) {
	c := gocrafty.DefaultConfig()
	var zero *gocrafty.ServerConfig

	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		data, err := json.MarshalIndent(c, "", "\t")

		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}

		if err := os.WriteFile("config.json", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}

		return &c, nil
	}

	data, err := os.ReadFile("config.json")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}

	if err := json.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}

	return &c, nil
}
