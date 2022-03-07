package config

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

func (config *Config) SetApplicationName(applicationName string) {
	config.applicationName = applicationName
}

func (config *Config) SetCompanyName(companyName string) {
	config.companyName = companyName
}
