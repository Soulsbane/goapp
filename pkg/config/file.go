package config

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml/v2"
)

func (config Config) OpenConfigFile() {
	fileName, _ := config.GetUserConfigFilePath()
	data, _ := ioutil.ReadFile(fileName)
	err := toml.Unmarshal(data, &config.Values)

	if err != nil {
		panic(err)
	}
}

func (config Config) SaveConfigFile() {
	fileName, _ := config.GetUserConfigFilePath()
	data, err := toml.Marshal(&config.Values)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(fileName, data, 0666)
}
