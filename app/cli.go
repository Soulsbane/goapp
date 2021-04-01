package goapp

import (
	"github.com/alexflint/go-arg"
)


type GoApp struct {
	Name string
	Version string
}

// NewCmdLineApp returns a new CLI instance with sensible defaults
func NewCmdLineApp(name string, version string, input interface{})*GoApp {
	var app GoApp
	app.Name = name
	app.Version = version

	arg.MustParse(input)
	return &app
}
