package cli

import (
	"github.com/alexflint/go-arg"
	"github.com/pterm/pterm"
)

type GoApp struct {
	Name    string
	Version string
	Debug   bool
}

// NewCmdLineApp returns a new CLI instance with sensible defaults
//func NewCmdLineApp(name string, version string, args interface{}) *GoApp {
func NewCmdLineApp(name string, version string, args interface{}) *GoApp {
	var app GoApp
	app.Name = name
	app.Version = version

	arg.MustParse(args)
	return &app
}

// PrintError print a message in red
func (app GoApp) PrintError(msg string) {
	pterm.Error.Println(msg)
}

// PrintWarning print a message in yellow
func (app GoApp) PrintWarning(msg string) {
	pterm.Warning.Println(msg)
}
