// package cmd implements methods for automating commandline tasks.
package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// RunLoud executes a command string using default file descriptors
func RunLoud(command string) error {
	cmd := Parse(command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// Parse parses a string into a cmd struct
func Parse(command string) *exec.Cmd {
	cmdlist := strings.Split(command, " ")
	return exec.Command(cmdlist[0], cmdlist[1:]...)
}

// RunSilent runs a comand string silently
func RunSilent(command string) error {
	cmd := Parse(command)
	return cmd.Run()
}
