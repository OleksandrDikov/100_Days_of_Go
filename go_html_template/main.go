package main

import (
	"html/template"
	"os"
)

type Inventory struct {
	AlertName      string
	UrlToDashboard string
}

func main() {
	message := Inventory{"alert", "http://test.local"}
	tmpl, err := template.New("test").Parse("{{.UrlToDashboard}} items are made of {{.AlertName}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, message)
	if err != nil {
		panic(err)
	}
}
