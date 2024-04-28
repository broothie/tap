package main

import (
	_ "embed"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

//go:embed README.md.sprintf
var readmeFormat string

func main() {
	cmd := exec.Command("tap", "-h")

	usage, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("").Parse(readmeFormat)
	if err != nil {
		panic(err)
	}

	readme, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := readme.Close(); err != nil {
			panic(err)
		}
	}()

	if err := tmpl.Execute(readme, map[string]string{"usage": strings.TrimSpace(string(usage))}); err != nil {
		panic(err)
	}
}
