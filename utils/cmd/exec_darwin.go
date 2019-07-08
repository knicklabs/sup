// +build darwin

package cmd

import (
	"os/exec"
)

var copyCmd = exec.Command("pbcopy")

func open(input string) *exec.Cmd {
	return exec.Command("open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command("open", "-a", appName, input)
}
