package commands

import (
	"fmt"
	"math/rand"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// RandIntn is a function variable wrapping rand.Intn so tests can override it.
var RandIntn = rand.Intn

// CommandCatch attempts to catch the specified Pokémon and add it to the user's
// Pokédex.
func CommandCatch(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	p, err := pokeapi.FetchPokemon(url)
	if err != nil {
		return err
	}

	if cfg.CaughtPokemon == nil {
		cfg.CaughtPokemon = make(map[string]pokeapi.Pokemon)
	}

	if _, exists := cfg.CaughtPokemon[p.Name]; exists {
		fmt.Printf("%s is already caught!\n", p.Name)
		return nil
	}

	// Seed the random generator for real usage.
	if RandIntn(p.BaseExperience) < 50 {
		fmt.Printf("%s was caught!\n", p.Name)
		cfg.CaughtPokemon[p.Name] = p
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}

	return nil
}
