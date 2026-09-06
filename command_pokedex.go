package main

import(
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokedex

	if len(pokedex) == 0 {
		fmt.Println("You have no Pokemon in your Pokedex. Start catching 'em all!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range pokedex {
			fmt.Printf("\t- %s\n", pokemon.Name) 
		}
	}

	return nil
}
