package main

// INFO Use go get github.com/Soulsbane/goapp@develop to get latest changes.
import (
	"fmt"

	"github.com/Soulsbane/goapp/pkg/cli"
	"github.com/Soulsbane/goapp/pkg/config"
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
		cli.WithName("OVER name"),
		cli.WithVendor("OVER vend"),
		cli.WithVersion("6.66"),
		cli.WithEnableDebug(true),
		cli.WithArgs(&args),
	)

	fmt.Println(app.Name)
	fmt.Println(args.NewApp.AppName)
	fmt.Println(args.Quiet)
	app.PrintError("This is an error")
	config.ConfigHello()
	fmt.Println(app.GetUserConfigDir())
}
