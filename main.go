package main

import (
	"time"
	"github.com/wNRG/pokedex/internal/pokeapi"
)

func main() {
	pokedex := map[string]pokeapi.Pokemon{}
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex: pokedex,
	}
	startRepl(cfg)
}

