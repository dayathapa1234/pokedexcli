package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Assign the real implementation to a function variable
var FetchLocationAreas = fetchLocationAreas

// The actual function (unexported)
func fetchLocationAreas(url string) (LocationAreaResponse, error) {
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
	return data, nil
}

func StringPtr(s string) *string {
	return &s
}