package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"rixz90/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
	var pokeDesc *pokeapi.PokemonDesc
	var err error

	url := fmt.Sprintf("%s/pokemon/%s", cfg.apiURL, pokemon)

	resp, ok := cfg.Cache.Get(url)
	if ok {
		err := json.Unmarshal(resp, &pokeDesc)
		if err != nil {
			return err
		}
	} else {
		pokeDesc, err = cfg.pokeapiClient.GetPokeDesc(url)
		if err != nil {
			return err
		}
	}

	chances := rand.Intn(300)
	baseExp := pokeDesc.BaseExperience
	name := pokeDesc.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if chances >= baseExp {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex[name] = *pokeDesc
	} else {
		fmt.Printf("%s escape!\n", name)
	}

	return nil
}
