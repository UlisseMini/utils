// package test implements easy ways to test
// the output of programs.
package test

import (
	"bytes"
	"fmt"
	"os/exec"
)

// a Case is used for testing the output of a program
type Case struct {
	// Path is the path to the executeable
	Path string
	// Expected is the expected output from the executeable
	Expected string
}

func (t Case) Run() error {
	// buffers for stdout and stderr
	var errbuf bytes.Buffer
	var outbuf bytes.Buffer

	cmd := &exec.Cmd{
		Path:   t.Path,
		Stderr: &errbuf,
		Stdout: &outbuf,
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	// convert them to strings
	stderr := string(bytes.TrimSpace(errbuf.Bytes()))
	stdout := string(bytes.TrimSpace(outbuf.Bytes()))
	if stderr != "" {
		return fmt.Errorf("stderr is not empty\nstderr: \n%s\n", stderr)
	}

	if stdout != t.Expected {
		return fmt.Errorf("Expected %q got %q", t.Expected, stdout)
	}
	return nil
}
