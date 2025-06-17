package pokeapi

import "github.com/dayathapa1234/pokedexcli/internal/pokecache"

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	Cache               *pokecache.Cache
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}
