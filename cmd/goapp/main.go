package main

// INFO Use go get github.com/Soulsbane/goapp@develop to get latest changes.
import (
	"fmt"
	"github.com/Soulsbane/goapp/cli"
	"github.com/Soulsbane/goapp/filelogger"
	"log/slog"
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
		&args,
		cli.WithName("GoApp"),
		cli.WithCompany("Raijinsoft"),
		cli.WithVersion("6.66"),
		cli.WithDebugMode(true),
	)

	//err := app.OpenConfigFile(&cfg)
	logger, _ := app.CreateFileLogger(nil)
	logger.Info("This is a test")
	slog.Info(filelogger.CreateLogFileName())

	//if err != nil && errors.Is(err, os.ErrNotExist) {
	//	app.PrintFatal("Config file not found.")
	//}

	switch {
	case args.NewApp != nil:
		createTemplate(args.NewApp.AppName)
		fmt.Println(args.Quiet)

	default:
		fmt.Println("Use -h to see help")
	}
}
