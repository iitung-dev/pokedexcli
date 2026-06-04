package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (ResponseLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	//check the cache
	data, ok := c.cache.Get(url)
	if ok {
		// cache hit 
		fmt.Println("cache hit!!")
		responseLocation := ResponseLocation{}
		err := json.Unmarshal(data, &responseLocation)
		if err != nil {
			return ResponseLocation{}, fmt.Errorf("Json unmarshal failed: %v", err)
		}

		return responseLocation, nil
	}
	fmt.Println("cache missed!!")
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocation{}, fmt.Errorf("Failed to construct GET request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocation{}, fmt.Errorf("GET request error: %v", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ResponseLocation{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocation{}, fmt.Errorf("Response body error: %v", err)
	}

	responseLocation := ResponseLocation{}
	err = json.Unmarshal(data, &responseLocation)
	if err != nil {
		return ResponseLocation{}, fmt.Errorf("Json unmarshal failed: %v", err)
	}
	
	c.cache.Add(url, data)

	return responseLocation, nil
}