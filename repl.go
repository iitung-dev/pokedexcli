package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iitung-dev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	previousLocationURL *string
	myPokemon map[string]pokeapi.ResponsePokemonInfo
}

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
		
		var arg string
		if len(words) > 1 {
			arg = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, arg)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error

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
			description: "Show the next list of location",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous list of location",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the pokemons in the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inpect a pokemon",
			callback:    commandInspect,
		},				
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}