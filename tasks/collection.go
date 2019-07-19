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
	pfn := fn
	files, err := ioutil.ReadDir(dir)
	// Even if we cannot read this dir, the program is still useful.
	// Proceed without previous file.
	if err != nil {
		return "", nil
	}

	match := -1
	for i, file := range files {
		if i > 0 && file.Name() == fn {
			match = i - 1
		}
	}

	if match >= 0 {
		prevFile := files[match]
		pfn = prevFile.Name()
	} else if len(files) > 0 {
		pfn = files[len(files)-1].Name()
	}

	if pfn == fn {
		return "", &fileNotFoundError{"Previous tasks not found"}
	}

	return pfn, nil
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
		prevFn = ""
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
		dat, err := ioutil.ReadFile(c.CurrFile)
		if err != nil {
			return "", err
		}
		return strings.Trim(string(dat), "\n"), nil
	} else if os.IsNotExist(err) {
		return "- No tasks today!", nil
	} else {
		return "", err
	}
}

// Previous returns the previous tasks.
func (c *Collection) previousTasks() (string, error) {
	if len(c.PrevFn) == 0 {
		return "- No previous tasks!", nil
	}

	if _, err := os.Stat(c.PrevFile); err == nil {
		dat, err := ioutil.ReadFile(c.PrevFile)
		if err != nil {
			return "", err
		}
		return strings.Trim(string(dat), "\n"), nil
	} else if os.IsNotExist(err) {
		return "- No previous tasks!", nil
	} else {
		return "", err
	}
}

// Add adds a task to the current file
func (c *Collection) Add(dat string) error {
	err := os.MkdirAll(c.Dir, 0755)
	if err != nil {
		return err
	}

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

	var title string
	if len(c.PrevFn) > 0 {
		title = fmt.Sprintf("Previously (%s)", dateFromFn(c.PrevFn))
	} else {
		title = "Previously"
	}

	return strings.Join([]string{
		title,
		dat,
	}, "\n"), nil
}
