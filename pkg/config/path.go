package config

import (
	"os"
	"path/filepath"
)

func (config Config) GetUserConfigDir() (string, error) {
	fileName, err := os.UserConfigDir()
	fileName = filepath.Join(fileName, config.companyName, config.applicationName)

	return fileName, err
}

func (config Config) GetConfigFilePath() (string, error) {
	fileName, err := config.GetUserConfigDir()
	fileName = filepath.Join(fileName, configFileName)

	return fileName, err
}
