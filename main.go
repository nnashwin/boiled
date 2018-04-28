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
	"path/filepath"
)

// Carton is a map of all of the eggs in your conf file
var Carton = struct {
	Eggs map[string]Egg `json:"eggs,omitempty"`
}{}

var createQs = []*survey.Question{
	{
		Name: "useCurrDir",
		Prompt: &survey.Confirm{
			Message: "Do you want to use the files in your curr dir as your boilerplate?",
		},
	},
	{
		Name:   "description",
		Prompt: &survey.Input{Message: "Please enter a description for your boilerplate."},
	},
}

func main() {
	errCol := color.New(color.FgRed).SprintFunc()

	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	dirPath := ".boiled"

	credStr := filepath.Join(homeDir, dirPath, "carton.json")

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
				if pathfinder.DoesExist(credStr) == false {
					return fmt.Errorf(errCol("You currently do not have any eggs.  Add a boilerplate and run again!"))
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

				color.Magenta("List of boilerplates in your Carton:\n")
				for _, egg := range Carton.Eggs {
					color.Magenta(fmt.Sprintf("Egg Name: %s\nDescription: %s\nHas Data: %t\n\n", egg.Nick, egg.Description, egg.HasData))
				}

				return nil
			},
		},

		{
			Name:    "egg create",
			Aliases: []string{"ec"},
			Usage:   "creates a new egg",
			Action: func(c *cli.Context) error {
				// create the carton file if it doesn't exist
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

				// set up prompt
				answers := struct {
					UseCurrDir  bool
					Description string
				}{}

				err = survey.Ask(createQs, &answers)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				if answers.UseCurrDir == true {
					err = CopyDir(".", filepath.Join(homeDir, dirPath, eggNick), make(map[string]struct{}))
					if err != nil {
						return fmt.Errorf(errCol(err))
					}
				}

				Carton.Eggs[eggNick] = Egg{eggNick, answers.UseCurrDir, answers.Description}

				// recopy / write the carton
				b, err := json.Marshal(Carton)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				ioutil.WriteFile(credStr, b, os.ModePerm)

				color.Magenta("The boilerplate \"%s\" has been added to your carton.", eggNick)

				return nil
			},
		},

		{
			Name:    "egg delete",
			Aliases: []string{"ed"},
			Usage:   "deletes an existing egg and the stored directory used to store it",
			Action: func(c *cli.Context) error {
				eggNick := c.Args()[0]
				if eggNick == "" {
					return fmt.Errorf(errCol("You must enter the name of the egg to delete it"))
				}

				content, err := ioutil.ReadFile(credStr)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				if len(content) > 0 {
					err = json.Unmarshal(content, &Carton)
					if err != nil {
						return fmt.Errorf(errCol(err))
					}
				}

				if _, ok := Carton.Eggs[eggNick]; ok == false {
					return fmt.Errorf(errCol("The egg \"%s\" doesn't exist in the carton"), eggNick)
				}

				if Carton.Eggs[eggNick].HasData != false {
					os.RemoveAll(filepath.Join(homeDir, dirPath, eggNick))
				}

				delete(Carton.Eggs, eggNick)

				b, err := json.Marshal(Carton)
				if err != nil {
					return fmt.Errorf(errCol(err))
				}

				ioutil.WriteFile(credStr, b, os.ModePerm)

				color.Magenta("The boilerplate \"%s\" has been deleted from your carton.", eggNick)

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
