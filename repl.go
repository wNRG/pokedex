package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, exists := cfg.commands[input[0]]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name	string
	description string
	callback func(*config) error
}


func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned 
}




