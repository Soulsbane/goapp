package main

import (
	"fmt"

	goapp "github.com/Soulsbane/goapp/app"
)

type MyCheckoutCmd struct {
	Branch string `arg:"positional"`
	Track  bool   `arg:"-t"`
}

var args struct {
	Checkout *MyCheckoutCmd `arg:"subcommand:checkout"`
	Quiet    bool           `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	app := goapp.NewCmdLineApp("Test App", "1.0", &args)
	fmt.Println(app.Name)
	fmt.Println(args.Checkout.Track)
}
