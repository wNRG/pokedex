package main

import (
	"strings"
	"io"
	"fmt"
	
	"github.com/chzyer/readline"
)

func startRepl(cfg *config) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "Pokedex > ",
		HistoryLimit: 100,
		InterruptPrompt: "^C",
		EOFPrompt: "exit",
	})

	if err != nil {
		fmt.Println("Error readline not working:", err)
		return
	}
	defer rl.Close()

	for {
		input, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(input) == 0 {
				return
			}
			continue
		}
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		cmd_input := cleanedInput[0]
		cmd_args := cleanedInput[1:]

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




