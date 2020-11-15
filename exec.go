package exec

import (
	"io"
	"io/ioutil"
	"os/exec"
)

// Env is environment of command line.
type Env struct {
	Dir    string
	Stdout io.Writer
	Stderr io.Writer
	err    error
}

// Run a command with the environment.
// You should check an error via Err method.
// If the environment hold an error, Run will not do anything.
func (e *Env) Run(name string, args ...string) {
	if e.err != nil {
		return
	}

	cmd := exec.Command(name, args...)
	cmd.Dir = e.Dir
	cmd.Stdout = ioutil.Discard
	if e.Stdout != nil {
		cmd.Stdout = e.Stdout
	}
	cmd.Stderr = ioutil.Discard
	if e.Stderr != nil {
		cmd.Stderr = e.Stderr
	}
	e.err = cmd.Run()
}

func (e *Env) Err() error {
	return e.err
}
