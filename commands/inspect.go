package commands

import (
	"fmt"

	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// CommandInspect prints details about a caught Pokemon. It does not perform any
// API calls, instead using the data stored on the configuration when the Pokemon
// was caught.
func CommandInspect(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	name := args[0]
	p, ok := cfg.CaughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")

	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
