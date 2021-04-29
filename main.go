package main

import (
	"fmt"

	"github.com/Soulsbane/goapp/pkg/cli"
)

type NewAppCommand struct {
	AppName string `arg:"positional"`
}

var args struct {
	NewApp *NewAppCommand `arg:"subcommand:init"`
	Quiet  bool           `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	app := cli.NewCmdLineApp("Test App", "1.0", &args)
	fmt.Println(app.Name)
	fmt.Println(args.NewApp.AppName)
	fmt.Println(args.Quiet)
	app.PrintError("This is an warning")
}
