package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/diegoalzate/pokedexcli/internal/cliOption"
	"github.com/diegoalzate/pokedexcli/internal/commands"
)

func StartRepl(config *cliOption.Config) {
	for {
		fmt.Print("pokedex> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		// get option
		command, args := cleanInput(text)
		opt, err := cliOption.GetOption(command, GetOptions(config, args))

		if err != nil {
			log.Fatal(err)
			continue
		}

		err = opt.Callback()

		if err != nil {
			log.Fatal(err)
			continue
		}
	}
}

func cleanInput(input string) (command string, args []string) {
	words := strings.Fields(input)
	lowerCommand := strings.ToLower(words[0])
	trimmed := strings.TrimSpace(lowerCommand)

	args = []string{}

	for _, word := range words[1:] {
		args = append(args, strings.ToLower(word))
	}

	return trimmed, args
}

func GetOptions(config *cliOption.Config, args []string) map[string]cliOption.Option {
	options := map[string]cliOption.Option{
		"exit": {
			Name:        "exit",
			Description: "exits program",
			Callback:    commands.ExitCommand,
		},
		"map": {
			Name:        "map",
			Description: "get next pokemon api location areas",
			Config:      config,
			Callback: func() error {
				return commands.Map(config)
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "get previous pokemon api location areas",
			Config:      config,
			Callback:    func() error { return commands.Mapb(config) },
		},
		"examine": {
			Name:        "examine",
			Description: "get more info from a location area",
			Config:      config,
			Callback:    func() error { return commands.Examine(config, args[0]) },
		},
		"catch": {
			Name:        "catch",
			Description: "catch a specific pokemon",
			Config:      config,
			Callback:    func() error { return commands.Catch(config, args[0]) },
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspect a specific pokemon",
			Config:      config,
			Callback:    func() error { return commands.Inspect(config, args[0]) },
		},
	}

	options["help"] = cliOption.Option{
		Name:        "help",
		Description: "displays help message",
		Callback:    func() error { return commands.HelpCommand(options) },
	}

	return options
}
