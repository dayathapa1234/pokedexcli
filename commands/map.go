package commands

import (
	"fmt"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

func CommandMap(cfg *pokeapi.Config) error {
	if cfg.NextLocationURL == nil {
		cfg.NextLocationURL = pokeapi.StringPtr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	}
	resp, err := pokeapi.FetchLocationAreas(*cfg.NextLocationURL)
	if err != nil {
		return err
	}
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous
	return nil
}

func CommandMapb(cfg *pokeapi.Config) error {
	if cfg.PreviousLocationURL == nil {
		fmt.Println("You're on the first page.");
		return nil;
	}

	resp, err := pokeapi.FetchLocationAreas(*cfg.PreviousLocationURL)
	if err != nil {
		return err
	}
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous
	return nil
}
