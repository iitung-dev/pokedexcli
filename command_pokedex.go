package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.myPokemon) == 0 {
		return fmt.Errorf("Your pokedex is empty... go catch some pokemons")
	}

	fmt.Println("Your pokedex:")
	for k, _ := range cfg.myPokemon {
		fmt.Printf(" - %v\n", k)
	}


	return nil
}