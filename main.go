package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "minikick"
	app.Usage = "Create and back some awesome projects"
	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "Create a new project! Just pass in a project name and target amount. (Dont use $ for the amount)",
			Action: func(c *cli.Context) {
				args := c.Args()
				fmt.Println(args[0], args[1])
			},
		},
		{
			Name:    "back",
			Aliases: []string{"b"},
			Usage:   "Back a project! The arguments are name, project name, credit card number, and an amount.",
			Action: func(c *cli.Context) {
				args := c.Args()
				fmt.Println(args[0], args[1], args[2], args[3])
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Display a project the backers and the amount they backed for a project!",
			Action: func(c *cli.Context) {
				fmt.Println(c.Args().First())
			},
		},
		{
			Name:    "backer",
			Aliases: []string{"br"},
			Usage:   "Display a list of projects that a backer has backed and the amounts backed",
			Action: func(c *cli.Context) {
				fmt.Println(c.Args().First())
			},
		},
	}
	app.RunAndExitOnError()
}
