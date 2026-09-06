package main

import(
	"fmt"
)
func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 { 
		return fmt.Errorf("Invalid Command: No input. Try `inspect <pokemon>`")
	}

	pokedex := cfg.pokedex
	name := args[0]

	pokemon, exists := pokedex[name]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range pokemon.Types {
			fmt.Printf("\t- %s\n", types.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
