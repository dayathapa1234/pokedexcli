package commands

import (
	"fmt"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// CommandExplore lists all Pokémon that can be encountered in the provided
// location area. It expects the area name as the first argument.
func CommandExplore(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: explore <area>")
	}

	area := args[0]
	fmt.Printf("Exploring %s...\n", area)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	resp, err := pokeapi.FetchLocationArea(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
