package main

import (
	"rixz90/internal/pokeapi"
	"rixz90/internal/pokecache"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		Cache:         *pokecache.NewCache(5 * time.Minute),
		apiURL:        "https://pokeapi.co/api/v2",
		Pokedex:       make(map[string]pokeapi.PokemonDesc),
	}

	startRepl(cfg)
}
