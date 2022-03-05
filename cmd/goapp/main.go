package main

// INFO Use go get github.com/Soulsbane/goapp@develop to get latest changes.
import (
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

func main() {
	app := cli.NewGoApp(
		cli.WithName("GoApp"),
		cli.WithCompany("Raijinsoft"),
		cli.WithVersion("6.66"),
		cli.WithDebugMode(true),
		cli.WithArgs(&args),
	)

	fmt.Println(app.GetName())
	fmt.Println(args.NewApp.AppName)
	fmt.Println(args.Quiet)
	app.PrintError("This is an error")
	app.PrintFatal("This is an fatal string")
	app.PrintWarning("This is an warning string")
	app.PrintInfo("This is an info string")
	app.PrintDebug("This is an debug string")
	logger := app.CreateFileLogger("test.log", os.O_TRUNC)
	logger.Println("This is a test")

	app.Println("[blue]This is a test [green]and another test")
	path, _ := app.GetUserConfigDir()
	app.PrintInfo(path)
}
