package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"os"
)

var Carton = struct {
	Eggs map[string]Egg `json:"eggs,omitempty"`
}{}

func main() {
	errCol := color.New(color.FgRed).SprintFunc()

	app := cli.NewApp()
	app.Name = "Boiled"
	app.Version = "0.1.0"
	app.Usage = "a tool to create and manage boilerplate applications"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		color.Magenta("Add a boilerplate project today!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "egg list",
			Aliases: []string{"el"},
			Usage:   "list all of the eggs you have in your boiled config",
			Action: func(c *cli.Context) error {
				homeDir, err := homedir.Dir()
				if err != nil {
					panic(err)
				}

				if pathfinder.DoesExist(homeDir+"/.boiled/eggCarton.json") == false {
					return fmt.Errorf(errCol("You currently do not have any eggs.  Add a boilerplate and run again!"))
				}

				return nil
			},
		},

		{
			Name:    "carton create",
			Aliases: []string{"cc"},
			Usage:   "creates a new carton for your eggs if it does not already exist",
			Action: func(c *cli.Context) error {
				homeDir, err := homedir.Dir()
				if err != nil {
					panic(err)
				}

				if pathfinder.DoesExist(homeDir+"./boiled/eggCarton.json") == true {
					return fmt.Errorf(errCol("You already have a carton with eggs.  Use the egg create command to add a new egg."))
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
