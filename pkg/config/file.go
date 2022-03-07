package config

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml/v2"
)

func (config Config) OpenConfigFile() {
	fileName, _ := config.GetConfigFilePath()
	data, _ := ioutil.ReadFile(fileName)
	err := toml.Unmarshal(data, &config.Values)

	if err != nil {
		panic(err)
	}
}
