// package test implements an easy way to test the output of programs.
package test

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

// a Case is used for testing the output of a program
type Case struct {
	// Path is the path to the executeable
	Path string

	// Expected is the expected output from the executeable
	Expected string

	// Time to wait for the program to finish
	Timeout time.Duration
}

func (t Case) Run() error {
	// context for timing out the command
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, t.Timeout)

	// create the command
	cmd := exec.CommandContext(ctx, t.Path)

	stdout, err := cmd.Output()
	if err != nil {
		if err.Error() == "signal: killed" {
			return errors.New("timeout exceeded")
		}

		return err
	}

	// check standard out
	if string(stdout) != t.Expected {
		return fmt.Errorf("Expected %q got %q", t.Expected, stdout)
	}

	return nil
}
