package cli

type GoAppOption func(*GoApp)

func WithName(name string) GoAppOption {
	return func(app *GoApp) {
		app.Name = name
	}
}

func WithCompany(company string) GoAppOption {
	return func(app *GoApp) {
		app.Company = company
	}
}

func WithVersion(version string) GoAppOption {
	return func(app *GoApp) {
		app.Version = version
	}
}

func WithEnableDebug(enable bool) GoAppOption {
	return func(app *GoApp) {
		app.EnableDebug = enable
	}
}

func WithArgs(args interface{}) GoAppOption {
	return func(app *GoApp) {
		app.Args = args
	}
}
