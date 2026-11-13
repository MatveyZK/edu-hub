package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/MatveyZK/edu-hub/internal/handlers"
	"github.com/MatveyZK/edu-hub/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "internal/configs/server.toml", "path to config file")
}

type Servicies struct{} // temporary solution tp init handlers

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(config)
	logger := s.Logger()

	handlers := handlers.NewHandlers(logger, Servicies{})
	s.SetRegistrar(handlers)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
