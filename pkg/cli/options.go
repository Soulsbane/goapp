package cli

type GoAppOption func(*GoApp)

func WithName(name string) GoAppOption {
	return func(app *GoApp) {
		app.name = name
	}
}

func Withcompany(company string) GoAppOption {
	return func(app *GoApp) {
		app.company = company
	}
}

func Withversion(version string) GoAppOption {
	return func(app *GoApp) {
		app.version = version
	}
}

func WithdebugMode(enable bool) GoAppOption {
	return func(app *GoApp) {
		app.debugMode = enable
	}
}

func WithArgs(args interface{}) GoAppOption {
	return func(app *GoApp) {
		app.Args = args
	}
}
