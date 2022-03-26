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

	if errors.Is(err, os.ErrNotExist) {
		app.PrintFatal("Config file not found.")
	} else {
		fmt.Println("OpenConfigFile Error: ", err)
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

	fmt.Println("cfg.Name: ", cfg.Name)
	cfg.Name = "What is this"
	err = app.SaveConfigFile(&cfg)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	/*switch {
	case app.Args.(NewAppCommand).AppName != nil:
		fmt.Println(args)
		createTemplate()

	default:
		fmt.Println("Use -h to see help")
	}*/
}
