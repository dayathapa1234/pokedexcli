package pokeapi

// PokemonStat describes a single stat for a Pokemon, e.g. HP or attack.
type PokemonStat struct {
	BaseStat int           `json:"base_stat"`
	Stat     NamedResource `json:"stat"`
}

// PokemonType wraps the named type for a Pokemon, e.g. fire or flying.
type PokemonType struct {
	Type NamedResource `json:"type"`
}

// Pokemon represents the subset of fields we care about for the inspect command.
type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}
