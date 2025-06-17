package pokeapi

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int                   `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}
