# Pokedex CLI

Pokedex CLI is a lightweight command line application for exploring data from the [PokéAPI](https://pokeapi.co/). It lets you browse location areas, inspect which Pokémon appear in each area and even try to catch them. All API responses are cached locally to keep the interface snappy.

## Features

- Browse location areas with `map` and `mapb` commands
- Explore a specific location with `explore <area>`
- Attempt to catch Pokémon with `catch <name>`
- Built-in caching layer to minimize repeated API requests
- Small REPL style interface implemented in pure Go

## Requirements

Go 1.22 or newer is required. The project uses no third-party dependencies.

## Getting Started

Clone the repository and build the binary:

```bash
go build
```

You can also run the tool directly:

```bash
go run .
```

## Usage

Once started you will see the prompt `Pokedex >`. The following commands are available:

- `help` &mdash; list all commands
- `map` &mdash; show the next page of location areas from the API
- `mapb` &mdash; show the previous page
- `explore <area>` &mdash; list the Pokémon that can be encountered in `area`
- `catch <pokemon>` &mdash; try your luck at catching a Pokémon
- `exit` &mdash; quit the program

Pagination state is kept between commands so you can page forward and backward through the location list. Network responses are cached for a short period so repeated exploration of the same areas is quick.

Example session:

```text
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
Pokedex > explore canalave-city-area
Found Pokemon:
 - bibarel
 - machoke
 - tentacool
Pokedex > catch bibarel
Throwing a Pokeball at bibarel...
bibarel was caught!
```

## Running Tests

Execute all unit tests with:

```bash
go test ./...
```

The tests cover the REPL helpers, command behaviour and the in-memory cache.

## Repository Layout

- `main.go` &mdash; application entry point
- `repl.go` &mdash; REPL implementation and command registration
- `commands/` &mdash; individual command handlers
- `internal/pokeapi` &mdash; minimal client for the PokéAPI
- `internal/pokecache` &mdash; simple time-based cache used by the client

Enjoy catching ’em all!
