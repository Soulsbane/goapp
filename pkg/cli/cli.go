package cli

import (
	"os"
	"path/filepath"

	"github.com/Soulsbane/goapp/pkg/filelogger"
	"github.com/alexflint/go-arg"
)

type GoApp struct {
	name      string
	company   string
	version   string
	debugMode bool
	Args      interface{}
}

// NewGoApp returns a new GoApp instance with sensible defaults
func NewGoApp(options ...GoAppOption) *GoApp {
	var emptyArgs struct{}

	const (
		defaultName      = "New Go Application"
		defaultCompany   = "My company Name"
		defaultVersion   = "1.0"
		defaultDebugMode = false
	)

	app := &GoApp{
		name:      defaultName,
		company:   defaultCompany,
		version:   defaultVersion,
		debugMode: defaultDebugMode,
		Args:      &emptyArgs,
	}

	for _, option := range options {
		option(app)
	}

	arg.MustParse(app.Args)
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

func (app GoApp) CreateFileLogger(fileName string, flag int) filelogger.FileLogger {
	logger := filelogger.New(fileName, app.GetUserConfigDir(), flag)
	return logger
}

// GetUserConfigDir Get the path to user's config dir with application and optionally company appened
func (app GoApp) GetUserConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, app.company, app.name)
}
