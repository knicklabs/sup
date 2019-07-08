package cmd

import (
	"io"
	"os/exec"
)

// Copy copies the input to the clipboard.
func Copy(input string) {
	echoCmd := exec.Command("echo", "-n", input)

	reader, writer := io.Pipe()
	
	echoCmd.Stdout = writer
	copyCmd.Stdin = reader

	echoCmd.Start()
	copyCmd.Start()

	echoCmd.Wait()
	writer.Close()
	copyCmd.Wait()
}

// Open opens the URI using the OS default program or
// the program specified by the appName.
func Open(input, appName string) {
	var cmd *exec.Cmd

	if appName == "default" {
		cmd = open(input)
	} else {
		cmd = openWith(input, appName)
	}

	cmd.Run()
}
