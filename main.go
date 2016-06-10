package main

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	shlex "github.com/flynn/go-shlex"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("steps")
	viper.AddConfigPath("/etc/buildenv/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s", err)
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	if viper.GetBool("debug") == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

type Step interface {
	Do() error
}

type CommandStep struct {
	Command string
}

func (s *CommandStep) Cmd() (*exec.Cmd, error) {
	parts, err := shlex.Split(s.Command)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	return cmd, nil
}

func (s CommandStep) Do() error {
	cmd, err := s.Cmd()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: watch with logging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func BuildEnv(c *cli.Context) error {
	for i := range viper.GetStringSlice("Steps") {

		key := fmt.Sprintf("Steps[%d] ", i)
		log.Info("Step key: ", key)
		step := viper.Sub(key)
		log.Info("here")

		log.Info("Running step: ", step.GetString("Name"))

		var s Step

		if step.GetString("Command") != "" {
			s = CommandStep{Command: step.GetString("Command")}
		}

		s.Do()
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Action = BuildEnv
	app.Run(os.Args)
}
