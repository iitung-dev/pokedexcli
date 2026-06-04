package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {

	getLocationList, err := cfg.pokeapiClient.LocationList(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = getLocationList.Next
	cfg.previousLocationURL = getLocationList.Previous

	for _, location := range getLocationList.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("You are on the first page")
	}
	getLocationList, err := cfg.pokeapiClient.LocationList(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = getLocationList.Next
	cfg.previousLocationURL = getLocationList.Previous

	for _, location := range getLocationList.Results {
		fmt.Println(location.Name)
	}

	return nil
}