package pokecache_test

import (
	"fmt"
	"testing"
	"time"

	pokecache "github.com/diegoalzate/pokedexcli/internal/pokeCache"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("example"),
		},
		{
			key: "https://example.com/path",
			val: []byte("more data"),
		},
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

func TestReadLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("key", []byte("value"))

	if _, ok := cache.Get("key"); !ok {
		t.Error("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("key"); ok {
		t.Error("expected to not find key")
		return
	}
}
