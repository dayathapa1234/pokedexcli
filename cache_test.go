package main

// Unit tests for the pokecache in-memory caching layer. They run in the main
// package and import the pokecache implementation to test its exported
// functionality.

import (
	"fmt"
	"testing"
	"time"

	"github.com/dayathapa1234/pokedexcli/internal/pokecache"
)

// TestAddGet ensures that values added to the cache can be retrieved by key.
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{key: "https://example.com", val: []byte("testdata")},
		{key: "https://example.com/path", val: []byte("moretestdata")},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

// TestReapLoop verifies that entries older than the cache interval are removed
// by the background reaper.
func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected to not find key")
		return
	}
}
