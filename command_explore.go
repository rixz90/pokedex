package main

import (
	"encoding/json"
	"fmt"
	"rixz90/internal/pokeapi"
)

func commandExplore(cfg *config, area_name string) error {
	var locationDesc *pokeapi.MapLocationDesc
	var err error

	url := fmt.Sprintf("%s/location-area/%s", cfg.apiURL, area_name)

	resp, ok := cfg.Cache.Get(url)
	if ok {
		err := json.Unmarshal(resp, &locationDesc)
		if err != nil {
			return err
		}
	} else {
		locationDesc, err = cfg.pokeapiClient.GetMapDesc(url)
		if err != nil {
			return err
		}
	}

	pokemons := locationDesc.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, poke := range pokemons {
		fmt.Printf("-%s\n", poke.Pokemon.Name)
	}

	dat, err := json.Marshal(locationDesc)
	if err != nil {
		return err
	}

	cfg.Cache.Add(url, dat)

	return nil
}
