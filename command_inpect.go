package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("Please provide pokemon name")
	}

	pokemon := args[0]
	v, ok := cfg.myPokemon[pokemon]
	if !ok {
		return fmt.Errorf("%s is not in your pokedex...", pokemon)
	}
	
	fmt.Printf("Name: %s\n", v.Name)
	fmt.Printf("Height: %v\n", v.Height)
	fmt.Printf("Weight: %v\n", v.Weight)

	fmt.Println("Stats:")
	for _, v := range v.Stats {
		fmt.Printf(" -%v: %v\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range v.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil
}