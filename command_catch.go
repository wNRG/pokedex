package main

import(
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Invalid Command: No Pokemon found. Try `catch <Pokemon>`")
	}

	pokemonName := args[0]

	data, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name)

	if tryCatch(data.BaseExperience) {
		fmt.Printf("%s was caught!\n", data.Name)
		cfg.pokedex[data.Name] = data
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", data.Name)
	}

	return nil
}


func catchChance(baseExp int) float64 {
	const (
		minChance = 0.02
		maxChance = 0.65
		maxBaseExp = 635
	)

	expRatio := float64(baseExp) / maxBaseExp

	if expRatio < 0 {
		expRatio = 0
	}
	if expRatio > 1 {
		expRatio = 1
	}
	
	return maxChance - expRatio*(maxChance - minChance)
}


func tryCatch(baseExp int) bool {
	chance := catchChance(baseExp)
	return rand.Float64() < chance
}
