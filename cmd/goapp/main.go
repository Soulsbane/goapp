package main

// INFO Use go get github.com/Soulsbane/goapp@develop to get latest changes.
import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/goapp/pkg/cli"
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
	cfg.Name = "Default Application Name is What?"

	app := cli.NewGoApp(
		cli.WithName("GoApp"),
		cli.WithCompany("Raijinsoft"),
		cli.WithVersion("6.66"),
		cli.WithDebugMode(true),
		cli.WithArgs(&args),
	)

	err := app.OpenConfigFile(&cfg)

	if errors.Is(err, os.ErrNotExist) {
		app.PrintFatal("Config file not found.")
	} else {
		fmt.Println(err)
	}

	fmt.Println(app.GetName())
	fmt.Println(args.NewApp.AppName)
	fmt.Println(args.Quiet)
	logger, err := app.CreateFileLogger("test.log", os.O_TRUNC)

	if err != nil {
		fmt.Println("New Logger Error: ", err)
	}

	logger.Println("This is a test")

	app.Println("[blue]This is a test [green]and another test")
	path, _ := app.GetUserConfigDir()
	app.PrintInfo(path)
	err = app.SaveConfigFile()

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
