package main

import (
	"io/ioutil"
	"sync"

	log "github.com/Sirupsen/logrus"

	"github.com/ghodss/yaml"
)

type Step struct {
	Name     string
	Command  string
	Parallel bool
	Steps    []Step
}

func (s *Step) Debug() {
	log.Debugf("Name: %#v", s.Name)
	log.Debugf("Command: %#v", s.Command)
	log.Debugf("Parallel: %#v", s.Parallel)
	for i, step := range s.Steps {
		log.Debug("Steps: ", i)
		step.Debug()
	}
}

func (s *Step) DoParallel() error {
	var wg sync.WaitGroup

	errs := MultiError{}

	for _, step := range s.Steps {
		wg.Add(1)
		go func(s Step) {
			defer wg.Done()
			err := s.Do()
			if err != nil {
				errs.Errors = append(errs.Errors, err)
			}
		}(step)
	}

	wg.Wait()
	return errs.GetError()
}

func (s *Step) Do() error {
	if s.Parallel {
		return s.DoParallel()
	}

	if s.Command != "" {
		return DoCommand(s.Command)
	}
	return nil
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

func Build(steps []Step) error {
	for _, step := range steps {
		log.Info("Running: ", step.Name)
		err := step.Do()
		if err != nil {
			return err
		}
	}
	return nil
}
