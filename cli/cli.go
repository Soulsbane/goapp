package cli

import (
	"path/filepath"

	"github.com/Soulsbane/goapp/pkg/config"
	"github.com/Soulsbane/goapp/pkg/filelogger"
	"github.com/alexflint/go-arg"
	"github.com/hashicorp/go-multierror"
)

type GoApp struct {
	name      string
	company   string
	version   string
	debugMode bool
	config.Config
}

// NewGoApp returns a new GoApp instance with sensible defaults
func NewGoApp(args any, options ...GoAppOption) *GoApp {
	const (
		defaultName      = "New Go Application"
		defaultCompany   = "My company Name"
		defaultVersion   = "1.0"
		defaultDebugMode = false
	)

	app := &GoApp{
		name:      config.DefaultName,
		company:   config.DefaultCompany,
		version:   config.DefaultVersion,
		debugMode: defaultDebugMode,
	}

	for _, option := range options {
		option(app)
	}

	arg.MustParse(args)

	app.SetConfigOptions(
		config.WithApplicationName(app.name),
		config.WithCompanyName(app.company),
	)

	return app
}

func (app GoApp) GetName() string {
	return app.name
}

func (app GoApp) GetCompany() string {
	return app.company
}

func (app GoApp) GetVersion() string {
	return app.version
}

func (app GoApp) IsDebugModeEnabled() bool {
	return app.debugMode
}

func (app *GoApp) EnableDebugMode() {
	app.debugMode = true
}

func (app *GoApp) DisableDebugMode() {
	app.debugMode = false
}

func (app GoApp) CreateFileLogger(fileName string, flag int) (*filelogger.FileLogger, error) {
	var result *multierror.Error
	var logger *filelogger.FileLogger

	if dir, err := app.Config.GetUserConfigDir(); err == nil {
		if logger, err := filelogger.New(fileName, filepath.Join(dir, "logs"), flag); err == nil {
			return logger, err
		} else {
			result = multierror.Append(result, err)
		}
	} else {
		result = multierror.Append(result, err)
	}

	return logger, result
}
