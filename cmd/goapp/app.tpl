package main

func main() {
	app := cli.NewGoApp(
		cli.WithName("{{.Name}}"),
		cli.WithCompany("{{.Company}}"),
		cli.WithVersion("{{.Version}}"),
		cli.WithDebugMode({{.EnableDebugMode}}),
	)
}
