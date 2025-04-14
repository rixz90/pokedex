package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetMapDesc(url string) (*MapLocationDesc, error) {
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

	locationDesc := MapLocationDesc{}
	err = json.Unmarshal(dat, &locationDesc)
	if err != nil {
		return nil, err
	}

	return &locationDesc, nil
}
