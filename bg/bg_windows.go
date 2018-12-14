// +build windows

package bg

import (
	"io"
	"os/exec"
)

// Spawn is a wrapper for executing commands
// in the background.
func Spawn(cmd *exec.Cmd, c io.ReadWriter) error {
	cmd.SysProcAttr.HideWindow = true

	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Stdin = c

	return cmd.Run()
}
