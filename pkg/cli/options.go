package cli

type GoAppOption func(*GoApp)

func WithName(name string) GoAppOption {
	return func(app *GoApp) {
		app.name = name
	}
}

func WithCompany(company string) GoAppOption {
	return func(app *GoApp) {
		app.company = company
	}
}

func WithVersion(version string) GoAppOption {
	return func(app *GoApp) {
		app.version = version
	}
}

func WithDebugMode(enable bool) GoAppOption {
	return func(app *GoApp) {
		app.debugMode = enable
	}
}

func WithArgs(args interface{}) GoAppOption {
	return func(app *GoApp) {
		app.Args = args
	}
}
