// +build !windows,!darwin

package cmd

import (
	"fmt"
	"os/exec"
)

var copyCmd = exec.Command("xclip", "-selection", "clipboard")

func open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
