package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed app.tpl
var templatesDir embed.FS

type templateVariables struct {
	Name            string
	Company         string
	Version         string
	EnableDebugMode bool
}

func createTemplate(appName string) {
	data, err := templatesDir.ReadFile("app.tpl")

	if err != nil {
		fmt.Println(err)
	}

	templateVars := templateVariables{"GoApp", "Company", "1.0.0", true}
	tmpl, err := template.New("test").Parse(string(data))

	if err != nil {
		fmt.Println("Failed to read template app.tpl")
	}
	err = tmpl.Execute(os.Stdout, templateVars)
	if err != nil {
		fmt.Println("Failed to read template app.tpl")
	}
}
