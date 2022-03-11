package config

import (
	"io/ioutil"

	"github.com/hashicorp/go-multierror"
	toml "github.com/pelletier/go-toml/v2"
)

func (config Config) OpenConfigFile(values interface{}) error {
	var result *multierror.Error

	if fileName, err := config.GetUserConfigFilePath(); err == nil {
		if data, err := ioutil.ReadFile(fileName); err == nil {
			if err = toml.Unmarshal(data, values); err == nil {
				result = multierror.Append(result, err)
			} else {
				result = multierror.Append(result, err)
			}
		} else {
			result = multierror.Append(result, err)
		}
	} else {
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

func (config *Config) SaveConfigFile(values interface{}) error {
	var result *multierror.Error

	fileName, err := config.GetUserConfigFilePath()

	if err != nil {
		result = multierror.Append(result, err)
	}

	data, err := toml.Marshal(&values)

	if err != nil {
		result = multierror.Append(result, err)
	}

	ioutil.WriteFile(fileName, data, 0666)

	return result.ErrorOrNil()
}
