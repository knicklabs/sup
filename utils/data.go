package utils

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

// TodaysTasks returns Today's tasks.
func TodaysTasks() (string, error) {
	fn := GetCurrentFilename()

	dir, err := GetTasksDir()
	if err != nil {
		return "", err
	}

	title := fmt.Sprintf("Today (%s):", Today())
	var body string

	data, err := ioutil.ReadFile(path.Join(dir, fn))
	if err != nil {
		body = "No tasks today!"
	} else {
		body = strings.Trim(string(data), "\n")
	}
	
	return strings.Join([]string{title, body}, "\n"), nil
}

// PreviousTasks returns the Previous Day's tasks.
func PreviousTasks() (string, error) {
	dir, err := GetTasksDir()
	if err != nil {
		return "", err
	}

	fn, err := GetPreviousFilename(dir)
	if err != nil {
		return "", err
	}

	if len(fn) == 0 {
		return "Previously\nNo previous tasks!", nil
	}

	data, err := ioutil.ReadFile(path.Join(dir, fn))
	if err != nil {
		return "", err
	}

	title := fmt.Sprintf("Previously (%s)", DayFromFilename(fn))
	body := strings.Trim(string(data), "\n")

	return strings.Join([]string{title, body}, "\n"), nil
}

// TodayAndPreviousTasks returns Today's and the Previous day's
// tasks.
func TodayAndPreviousTasks() (string, error) {
	ttxt, err := TodaysTasks()
	if err != nil {
		return "", err
	}

	ptxt, err := PreviousTasks()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{ptxt, ttxt}, "\n\n"), nil
}
