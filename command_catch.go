package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Can't catch... no pokemon provided")
	}

	pokemon := args[0]

	pokemonInfo, err := cfg.pokeapiClient.PokemonInfo(pokemon)
	if err != nil {
		return fmt.Errorf("Pokemon info request failed: %s", err)
	}

	capBaseExp := 635

	if pokemonInfo.BaseExperience >= capBaseExp {
		return fmt.Errorf("Cannot catch %s too strong!!", pokemon)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	threshold := 1.0 - (float64(pokemonInfo.BaseExperience)/float64(capBaseExp))

	if rand.Float64() < threshold {
		cfg.myPokemon[pokemon] = pokemonInfo
		fmt.Printf("Caught %s!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}