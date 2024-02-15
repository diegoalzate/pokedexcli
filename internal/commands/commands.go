package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/diegoalzate/pokedexcli/internal/cliOption"
)

func HelpCommand(options map[string]cliOption.Option) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n")
	for _, c := range options {
		fmt.Printf("%v: %v \n", c.Name, c.Description)
	}
	return nil
}

func ExitCommand() error {
	os.Exit(0)
	return nil
}

func Map(config *cliOption.Config) error {
	body, err := config.Client.GetLocationAreas(config.Next)

	if err != nil {
		return err
	}

	config.Update(body.Previous, body.Next)
	body.Print()
	return nil
}

func Mapb(config *cliOption.Config) error {
	if config.Previous == nil {
		return errors.New("no previous page")
	}
	body, err := config.Client.GetLocationAreas(config.Previous)

	if err != nil {
		return err
	}

	config.Update(body.Previous, body.Next)
	body.Print()
	return nil
}

func Examine(config *cliOption.Config, arg string) error {
	body, err := config.Client.GetLocationArea(arg)

	if err != nil {
		return err
	}
	body.Print()
	return nil
}

func Catch(arg string)
