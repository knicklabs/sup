package tasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/kyokomi/emoji.v1"

	"github.com/knicklabs/sup/config"
)

const (
	dateFmt = "2006-01-02"
	fileExt = "md"
	notFound = "- No tasks!"
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

func generateOutput(title, body string) string {
	return strings.Join([]string{
		title,
		body,
	}, "\n")
}

func generateTitle(base, fn string) string {
	if len(fn) > 0 {
		return fmt.Sprintf("%s (%s)", base, dateFromFn(fn))
	}
	return base
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

func (c *Collection) fileToString(file string) (string, error) {
	if (file == c.Dir) {
		return notFound, nil
	}

	if _, err := os.Stat(file); err == nil {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		return emoji.Sprint(strings.Trim(string(data), "\n")), nil
	} else if os.IsNotExist(err) {
		return notFound, nil
	} else {
		return "", err
	}
}

func (c *Collection) currentTasks() (string, error) {
	return c.fileToString(c.CurrFile)
}

func (c *Collection) dateTasks(d string) (string, error) {
	fn := path.Join(c.Dir, fmt.Sprintf("%s.%s", d, fileExt))
	return c.fileToString(fn)
}

func (c *Collection) previousTasks() (string, error) {
	return c.fileToString(c.PrevFile)
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
	body, err := c.currentTasks()
	if err != nil {
		return "", err
	}
	
	title := generateTitle("Current", c.CurrFn)
	return generateOutput(title, body), nil
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

// Date returns the date's tasks
func (c *Collection) Date(d string) (string, error) {
	body, err := c.dateTasks(d)
	if err != nil {
		return "", err
	}
	title := generateTitle("Date", d)
	return generateOutput(title, body), nil
}

// Previous returns the previous tasks
func (c *Collection) Previous() (string, error) {
	body, err := c.previousTasks()
	if err != nil {
		return "", err
	}

	title := generateTitle("Previously", c.PrevFn)
	return generateOutput(title, body), nil
}
