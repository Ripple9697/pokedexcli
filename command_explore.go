package main

import (
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("you didnt specify the location area")
	}
	locationArea := params[1]
	pokemon, err := cfg.pokeapiClient.ListPokemon(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s \n", pokemon.Pokemon.Name)
	}
	return nil

}
