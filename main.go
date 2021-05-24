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
	app := cli.NewCmdLineApp("Test App", "1.0", &args)
	fmt.Println(app.Name)
	fmt.Println(args.NewApp.AppName)
	fmt.Println(args.Quiet)
	app.PrintError("This is an error")
	config.ConfigHello()
}
