package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"io/ioutil"
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

				content, err := ioutil.ReadFile(credStr)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				// if the file contains a json object,  marshal it to the to the Carton
				if len(content) > 0 {
					err = json.Unmarshal(content, &Carton)
					if err != nil {
						return fmt.Errorf(errCol(err))
					}
				}

				// if the Carton is empty, create a carton
				if Carton.Eggs == nil {
					Carton.Eggs = make(map[string]Egg)
				}

				eggNick := c.Args()[0]
				if Carton.Eggs[eggNick] != (Egg{}) {
					return fmt.Errorf(errCol("The egg %s already exists."), eggNick)
				}

				useCurrDir := false
				prompt := &survey.Confirm{
					Message: "Do you want to use the files in your curr dir as your boilerplate?",
				}

				survey.AskOne(prompt, &useCurrDir, nil)

				Carton.Eggs[eggNick] = Egg{eggNick}

				// recopy / write the carton
				b, err := json.Marshal(Carton)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				ioutil.WriteFile(credStr, b, os.ModePerm)

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
