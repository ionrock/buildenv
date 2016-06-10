package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestLoadSteps(t *testing.T) {
	var data = `[
{"Name": "Do it",
 "Command": "echo \"Doing it!\""},

{"Name": "build backend",
 "Parallel": true}]
`
	steps := make([]Step, 0)

	err := json.Unmarshal([]byte(data), &steps)
	if err != nil {
		t.Error(err)
	}

	log.Infof("%#v", steps)
}

func TestLoadStepsFromFile(t *testing.T) {
	steps := LoadSteps("steps.yml")
	log.Infof("%#v", steps)
	if len(steps) != 2 {
		t.Fail()
	}
}
