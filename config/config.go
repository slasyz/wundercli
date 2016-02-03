package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const (
	configDirectoryName = "wundercli"
	configFileName      = "config.json"
)

var Config struct {
	AccessToken string
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

// Saves current config to file.
func SaveConfig() (err error) {
	configPath := getConfigPath()

	data, err := json.MarshalIndent(Config, "", "    ")
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

// Does all config-related: gets its path, opens it and puts it to variable.
// If config was opened successfully, return true,
// if it doesn't exist, return false,
// return an error otherwise.
func OpenConfig() (exists bool, err error) {
	configPath := getConfigPath()

	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, errors.New("config reading error")
		}
	} else {
		// Read config from file.
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			return false, err
		}

		// Decode it.
		err = json.Unmarshal(data, &Config)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
