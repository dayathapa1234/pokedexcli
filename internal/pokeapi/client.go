package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dayathapa1234/pokedexcli/internal/pokecache"
)

// Assign the real implementation to a function variable
var (
	FetchLocationAreas = fetchLocationAreas
	FetchLocationArea  = fetchLocationArea
	Cache              *pokecache.Cache
)

// The actual function (unexported)
func fetchLocationAreas(url string) (LocationAreaResponse, error) {
	if Cache != nil {
		if data, ok := Cache.Get(url); ok {
			var cached LocationAreaResponse
			if err := json.Unmarshal(data, &cached); err == nil {
				return cached, nil
			}
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("bad status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("read error: %w", err)
	}

	var data LocationAreaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	if Cache != nil {
		Cache.Add(url, body)
	}

	return data, nil
}

// fetchLocationArea retrieves details for a single location area. The response
// is cached in the global Cache so repeated calls for the same URL are fast.
func fetchLocationArea(url string) (LocationAreaExploreResponse, error) {
	if Cache != nil {
		if data, ok := Cache.Get(url); ok {
			var cached LocationAreaExploreResponse
			if err := json.Unmarshal(data, &cached); err == nil {
				return cached, nil
			}
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaExploreResponse{}, fmt.Errorf("request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaExploreResponse{}, fmt.Errorf("bad status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaExploreResponse{}, fmt.Errorf("read error: %w", err)
	}

	var data LocationAreaExploreResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreaExploreResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	if Cache != nil {
		Cache.Add(url, body)
	}

	return data, nil
}

func StringPtr(s string) *string {
	return &s
}
