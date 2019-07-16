package tasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/knicklabs/sup/config"
)

const (
	dateFmt = "2006-01-02"
	fileExt = "md"
)

// Collection represents a collection of tasks
type Collection struct {
	Dir      string
	CurrFile string
	CurrFn   string
	PrevFile string
	PrevFn   string
}

type fileNotFoundError struct {
	msg string
}

func (e *fileNotFoundError) Error() string {
	return e.msg
}

func dateFromFn(fn string) string {
	return strings.Replace(fn,
		strings.Join([]string{".", fileExt}, ""), "", 1)
}

func getFilename(daysFromNow int) string {
	t := time.Now().AddDate(0, 0, -daysFromNow)
	return fmt.Sprintf("%s.%s", t.Format(dateFmt), fileExt)
}

func getPrevFilename(dir string, fn string) (string, error) {
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
	} else if len(files) > 0 {
		return files[len(files)-1].Name(), nil
	}

	return "", &fileNotFoundError{"Previous tasks not found"}
}

// NewCollection initializes a collection.
func NewCollection(cfg *config.Config) (*Collection, error) {
	dir, err := cfg.AbsoluteTasksPath()
	if err != nil {
		return nil, err
	}

	currFn := getFilename(0)
	prevFn, err := getPrevFilename(dir, currFn)
	if err != nil {
		if _, ok := err.(*fileNotFoundError); !ok {
			return nil, err
		}
	}

	return &Collection{
		Dir:      dir,
		CurrFile: path.Join(dir, currFn),
		CurrFn:   currFn,
		PrevFile: path.Join(dir, prevFn),
		PrevFn:   prevFn,
	}, nil
}

// Current returns the current tasks.
func (c *Collection) currentTasks() (string, error) {
	if _, err := os.Stat(c.CurrFile); err == nil {
		data, err := ioutil.ReadFile(c.CurrFile)
		return strings.Trim(string(data), "\n"), err
	} else if os.IsNotExist(err) {
		return "- No tasks today!", nil
	} else {
		return "", err
	}
}

// Previous returns the previous tasks.
func (c *Collection) previousTasks() (string, error) {
	if len(c.PrevFile) == 0 {
		return "- No previous tasks!", nil
	}

	dat, err := ioutil.ReadFile(c.PrevFile)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(dat), "\n"), err
}

// Add adds a task to the current file
func (c *Collection) Add(dat string) error {
	f, err := os.OpenFile(c.CurrFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(dat)
	return nil
}

// Current returns the current tasks
func (c *Collection) Current() (string, error) {
	dat, err := c.currentTasks()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{
		fmt.Sprintf("Today (%s)", dateFromFn(c.CurrFn)),
		dat,
	}, "\n"), nil
}

// CurrentAndPrevious returns the current and previous tasks.
func (c *Collection) CurrentAndPrevious() (string, error) {
	prev, err := c.Previous()
	if err != nil {
		return "", err
	}

	curr, err := c.Current()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{prev, curr}, "\n\n"), nil
}

// Previous returns the previous tasks
func (c *Collection) Previous() (string, error) {
	dat, err := c.previousTasks()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{
		fmt.Sprintf("Previously (%s)", dateFromFn(c.PrevFn)),
		dat,
	}, "\n"), nil
}
