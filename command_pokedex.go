package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params []string) error {
	if len(cfg.pokedex) < 1 {
		return fmt.Errorf("No caputured pokemon")
	}
	fmt.Println("Your Pokedex:")
	for k,_ := range cfg.pokedex {
		fmt.Printf(" - %s\n",k)
	}
	return nil
}
	