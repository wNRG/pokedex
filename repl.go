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

		cmd_input := input[0]
		cmd_args := input[1:]

		command, exists := cfg.commands[cmd_input]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg, cmd_args...); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name	string
	description string
	callback func(*config, ...string) error
}


func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned 
}




