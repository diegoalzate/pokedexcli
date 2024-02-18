package commands

import (
	"errors"
	"fmt"
	"os"

	"math/rand"

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

func Catch(config *cliOption.Config, arg string) error {
	body, err := config.Client.GetPokemon(arg)

	chanceToCatch := rand.Intn(body.BaseExperience)

	fmt.Printf("Throwing to catch %s...\n", body.Name)

	if chanceToCatch > 40 {
		fmt.Printf("%s escaped!\n", body.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", body.Name)

	config.Pokedex[body.Name] = *body

	if err != nil {
		return err
	}

	return nil
}

func Inspect(config *cliOption.Config, arg string) error {
	body, found := config.Pokedex[arg]

	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s \n", body.Name)
	fmt.Printf("Height: %v \n", body.Height)
	fmt.Printf("Weight: %v \n", body.Weight)
	fmt.Println("Stats:")

	for _, stat := range body.Stats {
		fmt.Printf("  -%v: %v \n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range body.Types {
		fmt.Printf("  -%v \n", t.Type.Name)
	}
	return nil
}
