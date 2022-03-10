package config

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

func (config *Config) SetConfigPath(options ...ConfigOption) {
	for _, option := range options {
		option(config)
	}
}
