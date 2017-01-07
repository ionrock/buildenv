package main

import (
	"encoding/json"
	"log"
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

	log.Printf("%#v", steps)
}

func TestLoadStepsFromFile(t *testing.T) {
	steps := LoadSteps("steps.yml")
	log.Printf("%#v", steps)
	if len(steps) != 2 {
		t.Fail()
	}
}
