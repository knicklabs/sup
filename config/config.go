package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	configDir     = ".sup"
	configFile    = "config.json"
	defaultDir    = "~/.sup"
	defaultEditor = "default"
	homeDir       = "~/"
	taskDir       = "tasks"
)

type configNotFoundError struct {
	msg string
}

func (e *configNotFoundError) Error() string {
	return e.msg
}

// Config represents CLI configuration.
type Config struct {
	Dir    string `json:"dir"`
	Editor string `json:"editor"`
}

func getBaseDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return normalizeOsPath(dir), nil
}

func getConfigDir() (string, error) {
	dir, err := getBaseDir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, configDir), nil
}

func loadConfigData(dir string) ([]byte, error) {
	p := path.Join(dir, configFile)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, &configNotFoundError{"File does not exist"}
	}

	f, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return []byte(f), nil
}

func newConfig() *Config {
	return &Config{
		Dir:    defaultDir,
		Editor: defaultEditor,
	}
}

func normalizeOsPath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

// Get gets the Config.
func Get() (*Config, error) {
	dir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	cfg := newConfig()
	data, err := loadConfigData(dir)
	if err != nil {
		if _, ok := err.(*configNotFoundError); ok {
			return cfg, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// AbsoluteTasksPath returns the absolute path to tasks.
func (c *Config) AbsoluteTasksPath() (string, error) {
	if !strings.HasPrefix(c.Dir, homeDir) {
		return path.Join(c.Dir, taskDir), nil
	}

	dir, err := getBaseDir()
	if err != nil {
		return "", err
	}

	abs := strings.Replace(c.Dir, homeDir,
		strings.Join([]string{dir, string(os.PathSeparator)}, ""), 1)
	return normalizeOsPath(path.Join(abs, taskDir)), nil
}

// Reset resets the Config
func (c *Config) Reset() error {
	c.Dir = defaultDir
	c.Editor = defaultEditor
	return c.Save()
}

// Save saves the Config
func (c *Config) Save() error {
	dir, err := getConfigDir()
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	p := path.Join(dir, configFile)
	return ioutil.WriteFile(p, file, 0644)
}
