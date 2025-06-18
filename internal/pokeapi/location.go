package pokeapi

import "github.com/dayathapa1234/pokedexcli/internal/pokecache"

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	Cache               *pokecache.Cache
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

// NamedResource represents a generic object from the API that exposes a name
// and URL. Only the name field is currently used by the client.
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonEncounter describes a single Pokémon encountered in a location area.
type PokemonEncounter struct {
	Pokemon NamedResource `json:"pokemon"`
}

// LocationAreaExploreResponse contains the subset of fields from the API used
// by the explore command. It lists all Pokémon that can be encountered in the
// specified area.
type LocationAreaExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
