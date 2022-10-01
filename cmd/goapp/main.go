package main

// INFO Use go get github.com/Soulsbane/goapp@develop to get latest changes.
import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/goapp/pkg/cli"
	"github.com/alexflint/go-arg"
)

type NewAppCommand struct {
	AppName string `arg:"positional"`
}

var args struct {
	NewApp *NewAppCommand `arg:"subcommand:init"`
	Quiet  bool           `arg:"-q"` // this flag is global to all subcommands
}

type MyConfig struct {
	Name string
}

var cfg MyConfig

func main() {
	app := cli.NewGoApp(
		cli.WithName("GoApp"),
		cli.WithCompany("Raijinsoft"),
		cli.WithVersion("6.66"),
		cli.WithDebugMode(true),
	)

	arg.MustParse(&args)
	err := app.OpenConfigFile(&cfg)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		app.PrintFatal("Config file not found.")
	}

	switch {
	case args.NewApp != nil:
		createTemplate(args.NewApp.AppName)

	default:
		fmt.Println("Use -h to see help")
	}
}
