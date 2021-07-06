package cli

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/pterm/pterm"
)

type GoApp struct {
	Name        string
	Vendor      string
	Version     string
	EnableDebug bool
	Args        interface{}
}

type GoAppOption func(*GoApp)

// NewGoApp returns a new GoApp instance with sensible defaults
func NewGoApp(options ...GoAppOption) *GoApp {
	var emptyArgs struct{}

	const (
		defaultName        = "New Go Application"
		defaultVendor      = "My Vendor Name"
		defaultVersion     = "1.0"
		defaultEnableDebug = false
	)

	app := &GoApp{
		Name:        defaultName,
		Vendor:      defaultVendor,
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

func WithName(name string) GoAppOption {
	return func(app *GoApp) {
		app.Name = name
	}
}

func WithVendor(vendor string) GoAppOption {
	return func(app *GoApp) {
		app.Vendor = vendor
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

// CreateLogger Creates a logger using the path of GetUserConfigDir for file
func (app GoApp) CreateLogger(flag int) *log.Logger {
	os.MkdirAll(app.GetUserConfigDir(), os.ModePerm)
	file, err := os.OpenFile(path.Join(app.GetUserConfigDir(), "log.txt"), flag|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return logger
}

// GetUserConfigDir Get the path to user's config dir with application and optionally vendor appened
func (app GoApp) GetUserConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, app.Vendor, app.Name)
}
