package pokeapi

// Pokemon represents a Pokémon with the fields used by the CLI.
type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
