package main

import (
	"bufio"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	shlex "github.com/flynn/go-shlex"
)

func splitCommand(command string) ([]string, error) {
	parts, err := shlex.Split(os.ExpandEnv(command))
	if err != nil {
		return nil, err
	}

	return parts, nil
}

func watchCommand(c *exec.Cmd) {
	o, err := c.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating stdout pipe: ", err)
	}

	e, err := c.StderrPipe()
	if err != nil {
		log.Fatal("Error creating stderr pipe: ", err)
	}

	stdout := bufio.NewScanner(o)
	stderr := bufio.NewScanner(e)
	go func() {
		for stdout.Scan() {
			log.Infof("[stdout %d] %s", c.Process.Pid, stdout.Text())
		}
	}()

	go func() {
		for stderr.Scan() {
			log.Infof("[stderr %d] %s", c.Process.Pid, stderr.Text())
		}
	}()
}

func DoCommand(c string) error {
	parts, err := splitCommand(c)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(parts[0], parts[1:]...)

	watchCommand(cmd)

	err = cmd.Start()
	if err != nil {
		log.Fatal("Error starting cmd: ", err)
	}

	return cmd.Wait()
}
