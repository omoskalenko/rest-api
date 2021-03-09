package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/omoskalenko/rest-api/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}