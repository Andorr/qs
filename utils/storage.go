package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const configFile = "config.json"

type Config struct {
	Cookie string `json:"cookie"`
	Subjects map[string]int `json:"subjects"`
	People map[string]int `json:"people"`
}

var config *Config

// GetOrCreateConfig ...
func GetOrCreateConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}

	storagePath := GetStoragePath()
	// Initialize file path
	err := os.MkdirAll(storagePath, os.ModeDir)
	if err != nil {
		return nil, err
	}

	// Read config if it exists
	fullPath := fmt.Sprintf("%s\\%s", storagePath, configFile)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// The file does not exist, create it
		err := ioutil.WriteFile(fullPath, []byte("{}"), os.ModeDir)
		if err != nil {
			return nil, err
		}
		config = createConfig()
		return config, nil // Config was recently created, return empty config
	} else if err != nil {
		// Is other issues with reading the given file
		return nil, err
	}

	// Read config file
	configBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	// Parse bytes
	config = createConfig()
	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	// Making sure subjects and people are not nil
	if config.Subjects == nil {
		config.Subjects = make(map[string]int)
	}
	if config.People == nil {
		config.People = make(map[string]int)
	}

	return config, nil
}

// SaveConfig ...
func SaveConfig(cfg *Config) error {
	config = cfg

	storagePath := GetStoragePath()
	// Initialize file path
	err := os.MkdirAll(storagePath, os.ModeDir)
	if err != nil {
		return err
	}

	// Read config if it exists
	fullPath := fmt.Sprintf("%s\\%s", storagePath, configFile)
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Write config file
	err = ioutil.WriteFile(fullPath, bytes, 0777)
	if err != nil {
		return err
	}

	return nil
}

func MustGetConfig() *Config {
	config, err := GetOrCreateConfig()
	if err != nil {
		log.Fatalf(ConfigError, err.Error())
	}
	return config
}

func MustSaveConfig(cfg *Config) {
	err := SaveConfig(cfg)
	if err != nil {
		log.Fatalf(UnableToSave, err.Error())
	}
}

func createConfig() *Config {
	return &Config{
		Cookie:   "",
		Subjects: make(map[string]int),
		People: make(map[string]int),
	}
}