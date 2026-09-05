package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	
	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	cfg.nextLocationAreaURL = data.Next
	cfg.previousLocationAreaURL = data.Previous

	return nil
} 

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		fmt.Println("you're on the first page")
		fmt.Println()
		return nil
	}

	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	fmt.Println()

	cfg.nextLocationAreaURL = data.Next
	cfg.previousLocationAreaURL = data.Previous

	return nil
}
