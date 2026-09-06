package main

import (
	"fmt"
)

var order = []string{
	"help",
	"map",
	"mapb",
	"explore",
	"catch",
	"inspect",
	"pokedex",
	"exit",
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	commands := cfg.commands

	for _, name := range order {
		if cmd, exists := commands[name]; exists {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
	}

	fmt.Println()
	return nil
}
