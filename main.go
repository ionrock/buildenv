package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("buildenv")
	viper.AddConfigPath("/etc/buildenv/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Error("Fatal error config file: %s", err)
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	if viper.GetBool("debug") == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func BuildEnv(c *cli.Context) error {
	log.Info("Loading: ", c.String("steps"))
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
