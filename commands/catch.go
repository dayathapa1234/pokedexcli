package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// RandIntn is used to generate random integers. It is defined as a variable so
// tests can substitute a deterministic implementation.
var RandIntn = rand.Intn

// CommandCatch attempts to catch the specified Pokemon. If successful, the
// Pokemon's details are stored on the configuration so they can be inspected
// later.
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

	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Second)

	catchRate := 100.0 / (float64(p.BaseExperience) + 100.0) * 100.0
	if float64(RandIntn(100)) >= catchRate {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	if cfg.CaughtPokemon == nil {
		cfg.CaughtPokemon = make(map[string]pokeapi.Pokemon)
	}
	cfg.CaughtPokemon[name] = p
	fmt.Printf("%s was caught!\n", name)
	return nil
}
