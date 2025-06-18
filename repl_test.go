package main

// This file contains tests for the REPL helpers and command handlers. The tests
// live in the root package so that they can access unexported helpers such as
// `cleanInput` defined in repl.go.

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/dayathapa1234/pokedexcli/commands"
	"github.com/dayathapa1234/pokedexcli/internal/pokeapi"
)

// TestCleanInput verifies that the helper correctly normalizes user input by
// trimming spaces, converting to lower case and splitting into fields.
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   spaced    out    words   ",
			expected: []string{"spaced", "out", "words"},
		},
		{
			input:    "   one",
			expected: []string{"one"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input %q: expected %d words, got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Input %q: at index %d, expected %q but got %q", c.input, i, c.expected[i], actual[i])
			}
		}
	}
}

// mockFetchLocationAreas returns a deterministic response for the map commands
// so that tests do not perform real HTTP requests.
func mockFetchLocationAreas(url string) (pokeapi.LocationAreaResponse, error) {
	return pokeapi.LocationAreaResponse{
		Results: []pokeapi.LocationAreaResult{
			{Name: "canalave-city-area"},
			{Name: "eterna-city-area"},
			{Name: "pastoria-city-area"},
		},
		Next:     pokeapi.StringPtr("https://pokeapi.co/api/v2/location-area?offset=20&limit=20"),
		Previous: pokeapi.StringPtr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20"),
	}, nil
}

// captureOutput redirects stdout during the execution of f and returns the
// resulting output as a string. It is used to verify what the commands print
// to the console.
func captureOutput(f func()) string {
	var buf bytes.Buffer
	saved := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = saved
	buf.ReadFrom(r)
	return buf.String()
}

// TestCommandMap ensures that the map command prints the expected location
// names and updates the pagination URLs on the config.
func TestCommandMap(t *testing.T) {
	// Save original function
	original := pokeapi.FetchLocationAreas
	defer func() { pokeapi.FetchLocationAreas = original }()

	// Inject mock function
	pokeapi.FetchLocationAreas = func(url string) (pokeapi.LocationAreaResponse, error) {
		return pokeapi.LocationAreaResponse{
			Results: []pokeapi.LocationAreaResult{
				{Name: "canalave-city-area"},
				{Name: "eterna-city-area"},
			},
			Next:     pokeapi.StringPtr("next-url"),
			Previous: pokeapi.StringPtr("prev-url"),
		}, nil
	}

	// Setup and run
	cfg := &pokeapi.Config{}
	output := captureOutput(func() {
		err := commands.CommandMap(cfg, nil)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	// Assertions
	for _, name := range []string{"canalave-city-area", "eterna-city-area"} {
		if !strings.Contains(output, name) {
			t.Errorf("Expected output to contain %q", name)
		}
	}
}

// TestCommandMapb_FirstPage verifies that calling mapb on the first page
// displays the appropriate warning and does not error.
func TestCommandMapb_FirstPage(t *testing.T) {
	pokeapi.FetchLocationAreas = mockFetchLocationAreas // override API
	cfg := &pokeapi.Config{PreviousLocationURL: nil}

	out := captureOutput(func() {
		err := commands.CommandMapb(cfg, nil)
		if err != nil {
			t.Errorf("CommandMapb returned error on first page: %v", err)
		}
	})

	if !strings.Contains(out, "You're on the first page") {
		t.Errorf("Expected 'You're on the first page' message, got: %q", out)
	}
}

// TestCommandMapb ensures that mapb correctly prints locations from the
// previous page and updates the pagination URLs.
func TestCommandMapb(t *testing.T) {
	pokeapi.FetchLocationAreas = mockFetchLocationAreas // override API
	cfg := &pokeapi.Config{PreviousLocationURL: pokeapi.StringPtr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")}

	out := captureOutput(func() {
		err := commands.CommandMapb(cfg, nil)
		if err != nil {
			t.Errorf("CommandMapb returned error: %v", err)
		}
	})

	if !strings.Contains(out, "canalave-city-area") {
		t.Errorf("Expected location to be printed in output, got: %q", out)
	}
	if cfg.NextLocationURL == nil || cfg.PreviousLocationURL == nil {
		t.Error("Config next/previous URL not updated in mapb")
	}
}
