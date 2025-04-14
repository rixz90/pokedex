package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"rixz90/internal/pokeapi"
)

func commandMapf(cfg *config, _ string) error {
	var locationsResp pokeapi.RespShallowLocations
	var err error

	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL == nil {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
	} else {
		dat, ok := cfg.Cache.Get(*cfg.nextLocationsURL)
		if ok {
			err := json.Unmarshal(dat, &locationsResp)
			if err != nil {
				return err
			}
		} else {
			locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
			if err != nil {
				return err
			}
		}
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	encodedJSON, err := json.Marshal(locationsResp)
	if err != nil {
		return err
	}

	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL == nil {
		cfg.Cache.Add(cfg.apiURL+"/location-area", encodedJSON)
	}

	if cfg.nextLocationsURL != nil {
		cfg.Cache.Add(*cfg.nextLocationsURL, encodedJSON)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	return nil
}

func commandMapb(cfg *config, _ string) error {
	var locationsResp pokeapi.RespShallowLocations
	var err error

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL == nil {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	} else {

		dat, ok := cfg.Cache.Get(*cfg.prevLocationsURL)
		if ok {
			locationsResp = pokeapi.RespShallowLocations{}
			err := json.Unmarshal(dat, &locationsResp)
			if err != nil {
				return err
			}
		} else {
			locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
			if err != nil {
				return err
			}
		}
	}

	encodedJSON, err := json.Marshal(locationsResp)
	if err != nil {
		return err
	}

	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL == nil {
		cfg.Cache.Add(cfg.apiURL, encodedJSON)
	}

	if cfg.prevLocationsURL != nil {
		cfg.Cache.Add(*cfg.prevLocationsURL, encodedJSON)
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	return nil
}
