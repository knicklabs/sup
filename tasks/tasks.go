package tasks

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/knicklabs/sup/utils"
)

// Current returns the current day's tasks.
func Current() (string, error) {
	f := utils.GetCurrentFilename()

	dir, err := utils.GetTasksDir()
	if err != nil {
		return "", err
	}

	title := fmt.Sprintf("Today (%s):", utils.Today())
	var body string

	txt, err := ioutil.ReadFile(path.Join(dir, f))
	if err != nil {
		body = "No tasks today!"
	} else {
		body = strings.Trim(string(txt), "\n")
	}

	return strings.Join([]string{title, body}, "\n"), nil
}

// Previous returns the previous day's tasks.
func Previous() (string, error) {
	dir, err := utils.GetTasksDir()
	if err != nil {
		return "", err
	}

	f, err := utils.GetPreviousFilename(dir)
	if err != nil {
		return "Previously:\nNo previous tasks!", nil
	}

	txt, err := ioutil.ReadFile(path.Join(dir, f))
	if err != nil {
		return "", err
	}

	title := fmt.Sprintf("Previously (%s):", utils.DayFromFilename(f))
	body := strings.Trim(string(txt), "\n")

	return strings.Join([]string{title, body}, "\n"), nil
}

// CurrentAndPrevious returns today's and the previous day's tasks.
func CurrentAndPrevious() (string, error) {
	ctxt, err := Current()
	if err != nil {
		return "", err
	}

	ptxt, err := Previous()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{ptxt, ctxt}, "\n\n"), nil
}
