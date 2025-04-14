package main

import "fmt"

func commandInspect(cfg *config, pokemon string) error {
	poke, ok := cfg.Pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println(poke.BaseExperience, poke.Name, poke.Stats)
	return nil
}

func pokedex(cfg *config, _ string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("No pokemon caught")
	}

	for k, _ := range cfg.Pokedex {
		fmt.Printf("-%s\n", k)
	}

	return nil
}
