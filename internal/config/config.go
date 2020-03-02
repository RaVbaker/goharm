package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

var configDirName = "goharm"

func getDefaultConfigDir() (string, error) {
	var configDirLocation string

	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "linux":
		// Use the XDG_CONFIG_HOME variable if it is set, otherwise
		// $HOME/.config/example
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			configDirLocation = xdgConfigHome
		} else {
			configDirLocation = filepath.Join(homeDir, ".config", configDirName)
		}

	default:
		configDirLocation = filepath.Join(homeDir, ".config", configDirName)
	}

	return configDirLocation, nil
}

type Config struct {
	General GeneralOptions
	Keys    map[string]map[string]string
}
type Authentication struct {
	AccessToken, UserRole, UserId string
}

type GeneralOptions struct {
	ApiHost        string
	Authentication Authentication
}

var DefaultConfig = Config{
	General: GeneralOptions{
		ApiHost: "https://api.rebased.harmonogram.rebased.pl",
	},
}

const configFileName = "config.toml"

func LoadConfig() (*Config, error) {
	configFile, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	conf := DefaultConfig

	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		return &conf, nil
	} else if err != nil {
		return nil, err
	}

	if _, err = toml.DecodeFile(configFile, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func getConfigPath() (string, error) {
	configDir, err := getDefaultConfigDir()
	if err != nil {
		return "", err
	}
	configFile := filepath.Join(configDir, configFileName)
	return configFile, nil
}

func UpdateConfig(config *Config) error {
	configFile, err := getConfigPath()
	if err != nil {
		return err
	}
	configDir, _ := getDefaultConfigDir()
	err = os.MkdirAll(configDir, 0750)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(configFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(config)
	return err
}
