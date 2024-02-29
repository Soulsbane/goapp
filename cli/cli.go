package cli

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Soulsbane/goapp/config"
	"github.com/alexflint/go-arg"
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
		defaultDebugMode = false
	)

	app := &GoApp{
		name:      config.DefaultApplicationName,
		company:   config.DefaultCompanyName,
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

func (app *GoApp) GetName() string {
	return app.name
}

func (app *GoApp) GetCompany() string {
	return app.company
}

func (app *GoApp) GetVersion() string {
	return app.version
}

func (app *GoApp) IsDebugModeEnabled() bool {
	return app.debugMode
}

func (app *GoApp) EnableDebugMode() {
	app.debugMode = true
}

func (app *GoApp) DisableDebugMode() {
	app.debugMode = false
}

func (app *GoApp) CreateFileLogger(fileName string, options *slog.HandlerOptions) (*slog.Logger, error) {
	if dir, err := app.GetUserConfigDir(); err == nil {
		logDirectory := filepath.Join(dir, "logs")

		err := os.MkdirAll(logDirectory, os.ModePerm)

		if err != nil {
			return nil, fmt.Errorf("failed to make log file directory: %w", err)
		}

		file, err := os.OpenFile(filepath.Join(logDirectory, fileName), os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}

		return slog.New(slog.NewJSONHandler(file, options)), nil

	} else {
		return nil, fmt.Errorf("failed to get user config directory: %w", err)
	}
}
