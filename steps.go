package main

import (
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	shlex "github.com/flynn/go-shlex"
	"github.com/ghodss/yaml"
)

type Step struct {
	Name     string
	Command  string
	Parallel bool
	Steps    []Step
}

func (s *Step) Cmd() (*exec.Cmd, error) {
	parts, err := shlex.Split(s.Command)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	return cmd, nil
}

func (s Step) Do() error {
	cmd, err := s.Cmd()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: watch with logging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func LoadSteps(path string) []Step {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	steps := []Step{}

	err = yaml.Unmarshal(b, &steps)
	if err != nil {
		log.Fatal(err)
	}

	return steps
}

func Build(steps []Step) {
	for _, step := range steps {
		log.Info("Running: ", step.Name)
		step.Do()
	}
}
