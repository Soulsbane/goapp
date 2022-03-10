package config

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml/v2"
)

func (config Config) OpenConfigFile(values interface{}) {
	fileName, _ := config.GetUserConfigFilePath()
	data, _ := ioutil.ReadFile(fileName)

	err := toml.Unmarshal(data, values)

	if err != nil {
		panic(err)
	}
}

func (config *Config) SaveConfigFile(values interface{}) {
	fileName, _ := config.GetUserConfigFilePath()
	data, err := toml.Marshal(&values)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(fileName, data, 0666)
}
