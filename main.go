package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os"
)

func main() {
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
		fmt.Println(chalk.Magenta.Color("Add a boilerplate project today!"))
		return nil
	}

	app.Run(os.Args)
}
