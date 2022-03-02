package config

type Config struct {
	applicationName string
	companyName     string
}

func New(name string, company string) *Config {
	return &Config{
		applicationName: name,
		companyName:     company,
	}
}
