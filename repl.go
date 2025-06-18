package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dayathapa1234/pokedexcli/commands"
	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
	"github.com/dayathapa1234/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, []string) error
}

var command = map[string]cliCommand{}

func init() {
	command["map"] = cliCommand{
		name:        "map",
		description: "Explore 20 Pokémon location areas",
		callback:    commands.CommandMap,
	}
	command["help"] = cliCommand{
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	}
	command["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokédex",
		callback:    commandExit,
	}
	command["mapb"] = cliCommand{
		name:        "mapb",
		description: "Go back to the previous 20 Pokémon location areas",
		callback:    commands.CommandMapb,
	}
	command["explore"] = cliCommand{
		name:        "explore",
		description: "List Pokémon in the specified area",
		callback:    commands.CommandExplore,
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	pokeapi.Cache = cache
	cfg := &pokeapi.Config{Cache: cache}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		cmd, ok := command[cmdName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}

		if err := cmd.callback(cfg, words[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func commandExit(cfg *pokeapi.Config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeapi.Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range command {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
