package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "github.com/diegoalzate/pokedexcli/internal/pokeCache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	Cache pokecache.Cache
	Http  http.Client
}

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: *pokecache.NewCache(cacheInterval),
		Http: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(next *string) (*LocationAreasResponse, error) {
	out := LocationAreasResponse{}
	url := baseUrl

	if next != nil {
		url = *next
	}

	if body, ok := c.Cache.Get(url); ok {
		err := json.Unmarshal(body, &out)

		if err != nil {
			return &out, err
		}

		return &out, nil
	}

	resp, err := c.Http.Get(url)
	if err != nil {
		return &out, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return &out, err
	}

	err = json.Unmarshal(body, &out)

	if err != nil {
		return &out, err
	}

	return &out, nil
}

func (c *Client) GetLocationArea(name string) (*LocationAreaResponse, error) {
	out := LocationAreaResponse{}

	requestUrl := baseUrl + "/location-area" + "/" + name

	if body, ok := c.Cache.Get(requestUrl); ok {
		err := json.Unmarshal(body, &out)

		if err != nil {
			return &out, err
		}

		return &out, nil
	}

	resp, err := c.Http.Get(requestUrl)

	if err != nil {
		return &out, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return &out, err
	}

	err = json.Unmarshal(body, &out)

	if err != nil {
		return &out, err
	}

	return &out, nil
}

func (l *LocationAreasResponse) Print() {
	for _, l := range l.Results {
		fmt.Println(l.Name)
	}
}
func (l *LocationAreaResponse) Print() {
	fmt.Println("Found Pokemon")
	for _, l := range l.PokemonEncounters {
		fmt.Println("- " + l.Pokemon.Name)
	}
}
