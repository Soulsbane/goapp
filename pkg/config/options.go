package config

type ConfigOption func(*Config)

func WithApplicationName(name string) ConfigOption {
	return func(app *Config) {
		app.applicationName = name
	}
}

func WithCompanyName(company string) ConfigOption {
	return func(app *Config) {
		app.companyName = company
	}
}

func WithValues(values interface{}) ConfigOption {
	return func(app *Config) {
		app.Values = values
	}
}
