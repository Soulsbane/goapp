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
		cli.WithCompany("OVER vend"),
		cli.WithVersion("6.66"),
		cli.WithEnableDebug(true),
		cli.WithArgs(&args),
	)

	fmt.Println(app.Name)
}
