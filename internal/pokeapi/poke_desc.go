package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokeDesc(url string) (*PokemonDesc, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pokemonDesc := PokemonDesc{}
	err = json.Unmarshal(dat, &pokemonDesc)
	if err != nil {
		return nil, err
	}

	return &pokemonDesc, nil
}
