package main

import (
	"time"

	"github.com/diegoalzate/pokedexcli/internal/cliOption"
	"github.com/diegoalzate/pokedexcli/internal/client"
	"github.com/diegoalzate/pokedexcli/internal/repl"
)

func main() {
	client := client.NewClient(time.Second*5, time.Minute*5)
	cgf := &cliOption.Config{
		Client: client,
	}
	repl.StartRepl(cgf)
}
