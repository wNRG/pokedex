package main

import (
	"fmt"
	"github.com/wNRG/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config, args ...string) error {
	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	
	printLocationAreas(data)

	if data.Next != nil {
		cfg.nextLocationAreaURL = *data.Next
	} else  {
		cfg.nextLocationAreaURL = ""
	}

	if data.Previous != nil {
		cfg.previousLocationAreaURL = *data.Previous
	} else {
		cfg.previousLocationAreaURL = ""
	}

	return nil
} 

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == "" {
		fmt.Println("you're on the first page")
		fmt.Println()
		return nil
	}

	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationAreaURL,)
	if err != nil {
		return err
	}

	printLocationAreas(data)

	if data.Next != nil {
		cfg.nextLocationAreaURL = *data.Next
	} else  {
		cfg.nextLocationAreaURL = ""
	}

	if data.Previous != nil {
		cfg.previousLocationAreaURL = *data.Previous
	} else {
		cfg.previousLocationAreaURL = ""
	}

	return nil
}

func printLocationAreas(data pokeapi.LocationAreaResponse) {
	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()
}



