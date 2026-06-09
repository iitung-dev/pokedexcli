package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(name string) (ResponsePokemonInfo, error) {

	if name == "" {
		return ResponsePokemonInfo{}, fmt.Errorf("Please provide pokemon name or id\n", )
	}

	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponsePokemonInfo{}, fmt.Errorf("Pokemon Info New GET request fail: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokemonInfo{}, fmt.Errorf("Pokemon Info GET request fail: %v", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ResponsePokemonInfo{}, fmt.Errorf("Pokemon Info bad status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponsePokemonInfo{}, fmt.Errorf("Pokemon Info cannot read body: %s", err)
	}

	responsePokemonInfo := ResponsePokemonInfo{}
	err = json.Unmarshal(data, &responsePokemonInfo)
	if err != nil {
		return ResponsePokemonInfo{}, fmt.Errorf("Pokemon Info json unmarshal failed: %s", err)
	}

	return responsePokemonInfo, nil
}