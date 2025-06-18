package commands

import (
	"fmt"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// CommandMap prints the next page of Pokémon location areas. It ignores any
// additional arguments.
func CommandMap(cfg *pokeapi.Config, _ []string) error {
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

// CommandMapb prints the previous page of location areas. It also ignores any
// additional arguments.
func CommandMapb(cfg *pokeapi.Config, _ []string) error {
	if cfg.PreviousLocationURL == nil {
		fmt.Println("You're on the first page.")
		return nil
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
