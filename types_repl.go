package main

import (
	"rixz90/internal/pokeapi"
	"rixz90/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	Cache            pokecache.Cache
	apiURL           string
	Pokedex          map[string]pokeapi.PokemonDesc
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
