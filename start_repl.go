package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var secondCmd string
		if len(words) > 1 {
			secondCmd = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, secondCmd)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the location names",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the location names backward",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all pokemons in that area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch that pokemon & store it.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect catch pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all catch pokemon.",
			callback:    pokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
