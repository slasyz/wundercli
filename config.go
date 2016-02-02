package main

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"errors"
	"path"
	"github.com/slasyz/wundercli/api"
)

const (
	configDirectoryName = "wundercli"
	configFileName = "config.json"
)

type Config struct {
	AccessToken string
}

var (
	config Config
)

func NewConfig() Config {
	// I may populate it in future with some default variables.
	return Config{}
}

type ConfigDoesNotExist error

// Get config file path.
func getConfigPath() string {
	configDir := os.Getenv("XDG_CONFIG_HOME")

	if configDir != "" {
		// $XDG_CONFIG_HOME/wundercli/config.json
		return filepath.Join(configDir, configDirectoryName, configFileName)
	}

	homeDir := os.Getenv("HOME")

	if homeDir != "" {
		// If $XDG_CONFIG_HOME is not set then return
		// $HOME/.config/wundercli/config.json
		return filepath.Join(homeDir, ".config", configDirectoryName, configFileName)
	}

	// if $HOME is not set then return
	// config.json located in directory containing executable file.
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(absPath, configFileName)
}

// Opens config file, decodes JSON,
// sets accessToken variable in api package.
func parseConfigFile(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = NewConfig()
	err = json.Unmarshal(data, &config)

	api.SetAccessToken(config.AccessToken)

	return
}

// Saves current config to file.
func saveConfigFile(configPath string) (err error) {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return errors.New("JSON encode error.")
	}

	err = os.MkdirAll(path.Dir(configPath), 0744)
	if err != nil {
		return errors.New("error while creating config directory")
	}

	err = ioutil.WriteFile(configPath, append(data, byte('\n')), 0600)
	if err != nil {
		return errors.New("error while creating config file")
	}

	return nil
}
