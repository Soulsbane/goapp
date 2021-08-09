package cli

import (
	"os"
	"path/filepath"

	"github.com/Soulsbane/goapp/pkg/filelogger"
	"github.com/alexflint/go-arg"
	"github.com/pterm/pterm"
)

type GoApp struct {
	Name        string
	Company     string
	Version     string
	EnableDebug bool
	Args        interface{}
}

// NewGoApp returns a new GoApp instance with sensible defaults
func NewGoApp(options ...GoAppOption) *GoApp {
	var emptyArgs struct{}

	const (
		defaultName        = "New Go Application"
		defaultCompany     = "My Company Name"
		defaultVersion     = "1.0"
		defaultEnableDebug = false
	)

	app := &GoApp{
		Name:        defaultName,
		Company:     defaultCompany,
		Version:     defaultVersion,
		EnableDebug: defaultEnableDebug,
		Args:        &emptyArgs,
	}

	for _, option := range options {
		option(app)
	}

	arg.MustParse(app.Args)
	return app
}

// PrintError print a message in red
func (app GoApp) PrintError(msg string) {
	pterm.Error.Println(msg)
}

// PrintFatal print a message in light red
func (app GoApp) PrintFatal(msg string) {
	pterm.Fatal.Println(msg)
}

// PrintWarning print a message in yellow
func (app GoApp) PrintWarning(msg string) {
	pterm.Warning.Println(msg)
}

// PrintDebug print a message in gray
func (app GoApp) PrintDebug(msg string) {
	pterm.Debug.Println(msg)
}

// PrintInfo print a message in cyan
func (app GoApp) PrintInfo(msg string) {
	pterm.Info.Println(msg)
}

func (app GoApp) CreateFileLogger(fileName string, flag int) filelogger.FileLogger {
	logger := filelogger.New(fileName, app.GetUserConfigDir(), flag)
	return logger
}

// GetUserConfigDir Get the path to user's config dir with application and optionally company appened
func (app GoApp) GetUserConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, app.Company, app.Name)
}
