package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	requesteddt := RespShallowLocations{}
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL

	}
	data, good := c.cache.Get(url)
	if !good {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return requesteddt, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return requesteddt, err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return requesteddt, err
		}

		c.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &requesteddt)
	if err != nil {
		return requesteddt, err
	}

	return requesteddt, nil
}

func (c *Client) ListPokemon(locationArea string) (RespLocationArea, error) {
	obj := RespLocationArea{}

	url := baseURL + "/location-area/" + locationArea

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return obj, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return obj, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return obj, err
		}

		c.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
