# pokedexcli
This project provides a small interactive command line interface for
exploring Pokémon location areas using the [PokéAPI](https://pokeapi.co/).
It is implemented in Go and exposes a simple REPL with a few commands for
paging through available locations.

## Getting Started

A recent version of Go (1.22 or later) is required. Clone the repository
and build or run the program directly:

```bash
go run .
```

## Available Commands

- `help` - print a list of commands with a short description
- `map`  - list the next 20 Pokémon location areas
- `mapb` - go back to the previous page of location areas
- `exit` - quit the application

The `map` and `mapb` commands fetch data from the PokéAPI, so they require
an internet connection. The REPL keeps track of your current page so that
`map` advances forward and `mapb` moves backward.

## Running Tests

Unit tests can be executed with:

```bash
go test ./...
```

## Project Layout

- `main.go`           – program entry point starting the REPL
- `repl.go`           – command loop implementation and command registry
- `commands/`         – individual command handlers
- `internal/pokeapi`  – minimal client used to query the PokéAPI
