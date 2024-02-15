package cliOption

import (
	"fmt"

	"github.com/diegoalzate/pokedexcli/internal/client"
)

type Config struct {
	Next     *string
	Previous *string
	Client   client.Client
}

type Option struct {
	Name        string
	Description string
	Callback    func() error
	Config      *Config
}

func GetOption(input string, options map[string]Option) (*Option, error) {
	selected, ok := options[input]

	if !ok {
		return nil, fmt.Errorf("%v is not supported", input)
	}

	return &selected, nil
}

func (config *Config) Update(prev *string, next *string) {
	config.Next = next
	config.Previous = prev
}
