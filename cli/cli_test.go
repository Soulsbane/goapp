package cli_test

import (
	"testing"
)

func TestCli(m *testing.T) {
	/*app := cli.NewGoApp(
		cli.WithName("OVER name"),
		cli.WithCompany("OVER vend"),
		cli.WithVersion("6.66"),
		cli.WithDebugMode(true),
	)*/

	/*err := app.OpenConfigFile(&cfg)

	if errors.Is(err, os.ErrNotExist) {
		app.PrintFatal("Config file not found.")
	} else {
		fmt.Println("OpenConfigFile Error: ", err)
	}
	fmt.Println(app.GetName())
	fmt.Println(app.GetName())
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
	}*/
}
