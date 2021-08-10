package cli_test

import (
	"fmt"
	"testing"

	"github.com/Soulsbane/goapp/pkg/cli"
)

type NewAppCommand struct {
	AppName string `arg:"positional"`
}

var args struct {
	NewApp *NewAppCommand `arg:"subcommand:init"`
	Quiet  bool           `arg:"-q"` // this flag is global to all subcommands
}

func TestCli(m *testing.T) {
	app := cli.NewGoApp(
		cli.WithName("OVER name"),
		cli.Withcompany("OVER vend"),
		cli.Withversion("6.66"),
		cli.WithdebugMode(true),
		cli.WithArgs(&args),
	)

	fmt.Println(app.GetName())
}
