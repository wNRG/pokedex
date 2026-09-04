package main

import (
	"github.com/wNRG/pokedex/internal/pokeapi"
)
func main() {
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(cfg)
}

