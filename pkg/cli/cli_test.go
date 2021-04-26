package cli_test

import (
	"fmt"
	"testing"

	"github.com/Soulsbane/goapp/pkg/cli"
)

type MyCheckoutCmd struct {
	Branch string `arg:"positional"`
	Track  bool   `arg:"-t"`
}

var args struct {
	Checkout *MyCheckoutCmd `arg:"subcommand:checkout"`
	Quiet    bool           `arg:"-q"` // this flag is global to all subcommands
}

func TestCli(m *testing.T) {
	app := cli.NewCmdLineApp("Test App", "1.0", &args)
	fmt.Println(app.Name)
	fmt.Println(args.Checkout.Track)
}
