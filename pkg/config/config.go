package config

import gap "github.com/muesli/go-app-paths"

const (
	configFileName = "config.toml"
)

type Config struct {
	applicationName string
	companyName     string
	scope           *gap.Scope
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

	app.scope = gap.NewVendorScope(gap.User, app.applicationName, app.companyName)

	return app
}

func (config Config) GetConfigFilePath() (string, error) {
	fileName, err := config.scope.ConfigPath(configFileName)
	return fileName, err
}
