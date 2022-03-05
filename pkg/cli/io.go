package cli

import (
	"runtime"
	"strconv"

	"github.com/mitchellh/colorstring"
)

const (
	callerInfoColor = "[green]"
)

// Println Print a message using mitchellh/colorstring tags
func (app GoApp) Println(msg ...interface{}) {
	var args string

	for _, arg := range msg {
		args += arg.(string)
	}

	colorized := colorstring.Color(args)
	colorstring.Println(colorized)
}

func (app GoApp) PrintCallerInfo() {
	_, fileName, line, _ := runtime.Caller(2)
	app.Println(callerInfoColor, "└(", fileName, ":", strconv.Itoa(line), ")")
}

// PrintError print a message in red
func (app GoApp) PrintError(msg string) {
	app.Println("[_red_][black]  ERROR  [_default_][light_red] " + msg)
	app.PrintCallerInfo()
}

// PrintFatal print a message in light red
func (app GoApp) PrintFatal(msg string) {
	app.Println("[_red_][black]  FATAL  [_default_][light_red] " + msg)
	app.PrintCallerInfo()
}

// PrintWarning print a message in yellow
func (app GoApp) PrintWarning(msg string) {
	app.Println("[_yellow_][black]  WARNING  [_default_][yellow] " + msg)
	app.PrintCallerInfo()
}

// PrintDebug print a message in gray
func (app GoApp) PrintDebug(msg string) {
	if app.IsDebugModeEnabled() {
		app.Println("[_light_red_][black]  DEBUG  [_default_][light_red] " + msg)
		app.PrintCallerInfo()
	}
}

// PrintInfo print a message in cyan
func (app GoApp) PrintInfo(msg string) {
	app.Println("[_cyan_][black]  INFO  [_default_][light_cyan] " + msg)
	app.PrintCallerInfo()
}
