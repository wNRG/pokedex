package main

import (
	"github.com/wNRG/pokedex/internal/pokeapi"
)

type config struct {
	commands map[string]cliCommand
	pokeapiClient *pokeapi.Client
	pokedex map[string]pokeapi.Pokemon
	nextLocationAreaURL *string
	previousLocationAreaURL *string
} 

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},

		"map": {
			name: "map",
			description: "Display names of 20 location areas in the Pokemon world",
			callback: commandMapf,
		},

		"mapb": {
			name: "mapb",
			description: "Display the previous page of location areas",
			callback: commandMapb,
		},

		"explore": {
			name: "explore <area_name>",
			description: "Display all Pokemon located in the area",
			callback: commandExplore,
		},
		
		"catch": {
			name: "catch <pokemon>",
			description: "Attempt to catch a Pokemon and add them to your Pokedex",
			callback: commandCatch,
		},
		
		"pokedex": {
			name: "pokedex",
			description: "Display all the Pokemon you have caught",
			callback: commandPokedex,
		},

		"inspect": {
			name: "inspect <pokemon>",
			description: "View details of a Pokemon you have caught",
			callback: commandInspect,
		},

		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

	}
}

