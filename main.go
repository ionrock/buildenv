package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func BuildEnv(c *cli.Context) error {
	log.Print("Loading: ", c.String("steps"))
	steps := LoadSteps(c.String("steps"))

	for _, step := range steps {
		step.Debug()
	}
	Build(steps)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Action = BuildEnv
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "steps, s",
		},
	}
	app.Run(os.Args)
}
