package main

import(
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Invalid Command: No location area found. Try `explore <location_area>`")
	}

	locationArea := args[0]

	data, err := cfg.pokeapiClient.GetPokemonInArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", data.Name)
	fmt.Println("Found Pokemon:")

	for _, encounter := range data.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}

