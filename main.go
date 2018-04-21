package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"os"
)

// Carton is a map of all of the eggs in your conf file
var Carton = struct {
	Eggs map[string]Egg `json:"eggs,omitempty"`
}{}

func main() {
	errCol := color.New(color.FgRed).SprintFunc()

	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	credStr := homeDir + "/.boiled/carton.json"

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
				if pathfinder.DoesExist(homeDir+"/.boiled/eggCarton.json") == false {
					return fmt.Errorf(errCol("You currently do not have any eggs.  Add a boilerplate and run again!"))
				}

				return nil
			},
		},

		{
			Name:    "egg create",
			Aliases: []string{"ec"},
			Usage:   "creates a new egg",
			Action: func(c *cli.Context) error {
				if pathfinder.DoesExist(credStr) == false {
					err := pathfinder.CreateFile(credStr)
					if err != nil {
						return fmt.Errorf(errCol(err))
					}
				}

				fmt.Printf("%+v", c.Args())

				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
