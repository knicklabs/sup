package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

var (
	dateFmt = "2006-01-02"
	fileExt = "md"
	rootDir = ".sup"
	taskDir = "tasks"
)

// GetCurrentFilename returns the filename for Today's tasks
func GetCurrentFilename() string {
	return getFilename(0)
}

// GetPreviousFilename returns the name of the file for the
// previous days tasks.
func GetPreviousFilename(dir string) (string, error) {
	fn := GetCurrentFilename()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	match := -1
	for i, file := range files {
		if i > 0 && file.Name() == fn {
			match = i - 1
		}
	}

	if match >= 0 {
		prevFile := files[match]
		return prevFile.Name(), nil
	}

	if len(files) > 0 {
		return files[len(files)-1].Name(), nil
	}

	return "", nil
}

// GetDir returns the default directory.
func GetDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fp := path.Join(home, rootDir)
	return strings.Replace(fp, "\\", "/", -1), nil
}

// GetTasksDir returns the default directory to save tasks to.
func GetTasksDir() (string, error) {
	dir, err := GetDir()
	if err != nil {
		return "", err
	}
	return path.Join(dir, taskDir), nil
}

// MakeDir makes a directory.
func MakeDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// WriteStringToFile writes a string to a file.
func WriteStringToFile(dir, fn, str string) error {
	f, err := os.OpenFile(path.Join(dir, fn), os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(str)
	return nil
}

func getFilename(daysFromNow int) string {
	t := time.Now().AddDate(0, 0, -daysFromNow)
	return fmt.Sprintf("%s.%s", t.Format(dateFmt), fileExt)
}
