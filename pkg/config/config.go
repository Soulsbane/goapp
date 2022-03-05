package config

import (
	"os"
	"path/filepath"
)

const (
	configFileName = "config.toml"
)

type Config struct {
	applicationName string
	companyName     string
}

func New(options ...ConfigOption) *Config {
	const (
		defaultName    = "GoApp"
		defaultCompany = "GoCompanyApp"
	)

	app := &Config{
		applicationName: defaultName,
		companyName:     defaultCompany,
	}

	for _, option := range options {
		option(app)
	}

	return app
}

func (config *Config) SetApplicationName(applicationName string) {
	config.applicationName = applicationName
}

func (config *Config) SetCompanyName(companyName string) {
	config.companyName = companyName
}

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
