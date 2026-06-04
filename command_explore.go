package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	getExploreList, err := cfg.pokeapiClient.ExploreList(name)
	if err != nil {
		return fmt.Errorf("GET explore list error: %s", err)
	}

	for _, area := range getExploreList.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}

	return nil
}