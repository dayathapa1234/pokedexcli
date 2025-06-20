package commands

import (
	"fmt"
	"sort"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// CommandPokedex prints all caught Pokémon names. It ignores additional args.
func CommandPokedex(cfg *pokeapi.Config, _ []string) error {
	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	names := make([]string, 0, len(cfg.CaughtPokemon))
	for name := range cfg.CaughtPokemon {
		names = append(names, name)
	}
	sort.Strings(names)
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
