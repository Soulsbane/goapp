package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

const (
	configFileName = "config.toml"
)

type Config struct {
	applicationName string
	companyName     string
	Values          interface{}
}

func New(options ...ConfigOption) *Config {
	var emptyValues struct{}

	const (
		defaultName    = "GoApp"
		defaultCompany = "GoCompanyApp"
	)

	app := &Config{
		applicationName: defaultName,
		companyName:     defaultCompany,
		Values:          &emptyValues,
	}

	for _, option := range options {
		option(app)
	}

	return app
}

func (config Config) OpenConfigFile() {
	fileName, _ := config.GetConfigFilePath()
	data, _ := ioutil.ReadFile(fileName)
	err := toml.Unmarshal(data, &config.Values)

	if err != nil {
		panic(err)
	}
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
