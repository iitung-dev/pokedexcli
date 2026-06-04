package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreList(locationIdName string) (ResponseExplore, error) {

	if locationIdName == "" {
		return ResponseExplore{}, fmt.Errorf("Location name / id is missing!")
	}

	url := baseURL + "/location-area/" + locationIdName

	// check the cache 
	data, ok := c.cache.Get(url)
	if ok {
		// cache hit!!
		fmt.Println("Cache hit!!")
		responseExplore := ResponseExplore{}
		err := json.Unmarshal(data, &responseExplore) 
		if err != nil {
			return ResponseExplore{}, fmt.Errorf("Explore json unmarshal failed: %s", err)
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseExplore{}, fmt.Errorf("Explore New GET request error: %s", err)
	}

	fmt.Printf("Exploring %s...\n", locationIdName)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseExplore{}, fmt.Errorf("GET request error: %s", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ResponseExplore{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return ResponseExplore{}, fmt.Errorf("Response body error: %s", err)
	}

	fmt.Println("Found Pokemon:")
	responseExplore := ResponseExplore{}
	err = json.Unmarshal(data, &responseExplore) 
	if err != nil {
		return ResponseExplore{}, fmt.Errorf("Explore json unmarshal failed: %s", err)
	}

	c.cache.Add(url, data)

	return responseExplore, nil
}