package utils

import (
	"strings"
	"time"
)

// Today returns Today's date as a string.
func Today() string {
	return time.Now().Format("2006-01-02")
}

// DayFromFilename extracts the date from filename
// and returns it.
func DayFromFilename(fn string) string {
	return strings.Replace(fn, ".md", "", 1)
}
