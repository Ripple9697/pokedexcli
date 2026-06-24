package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("you didnt specify the pokemon's name")
	}
	pName := params[0]
	pokemon, err := cfg.pokeapiClient.ListPokemonDetail(pName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pName)
	
	if rand.Intn(pokemon.BaseExperience) > pokemon.BaseExperience - 100 {
		cfg.pokedex[pName] = pokemon
		fmt.Printf("%s was caught!\n",pName)
		return nil
	}
	fmt.Printf("%s escaped!\n",pName)
	return nil

}
