package main

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/ghodss/yaml"
)

type Step struct {
	Name     string
	Command  string
	OnFail   string
	Retry    int
	Parallel bool
	Steps    []Step
}

func (s *Step) Debug() {
	Debugf("Name: %#v", s.Name)
	Debugf("Command: %#v", s.Command)
	Debugf("Parallel: %#v", s.Parallel)
	for i, step := range s.Steps {
		Debug("Steps: ", i)
		step.Debug()
	}
}

func (s *Step) DoStep() error {
	if s.Command != "" {
		return DoCommand(s.Command, s.Name)
	}

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
	if s.Retry == 0 {
		s.Retry = 1
	}
	var err error
	for retry := 0; retry <= s.Retry; retry++ {
		err = s.DoStep()
		if err == nil {
			break
		}
		if s.OnFail != "" {
			log.Print("Error in Step %s. Running on fail command: %s", s.Name, s.OnFail)
			err = DoCommand(s.OnFail, s.Name)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return err
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
		log.Print("Running: ", step.Name)
		err := step.Do()
		if err != nil {
			log.Print(err)
		}
	}
	return nil
}
