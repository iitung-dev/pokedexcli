package main

import (
	"time"

	"github.com/iitung-dev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config {
		pokeapiClient: pokeClient,
		myPokemon: make(map[string]pokeapi.ResponsePokemonInfo),
	}

	startRepl(cfg)
}