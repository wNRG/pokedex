package main

func main() {
	cfg := config{
		commands: map[string]cliCommand{
			"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},

			"exit": {
				name: "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
			},
		},
	}
	startRepl(&cfg)
}

