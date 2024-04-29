package main

import (
	"bytes"
	_ "embed"
	"os"
	"strings"
	"text/template"

	"github.com/broothie/tap/internal/cli"
)

//go:embed README.md.tmpl
var readmeFormat string

func main() {
	var usage bytes.Buffer
	cli.New().UsageWriter(&usage).Usage(nil)

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

	if err := tmpl.Execute(readme, map[string]string{"usage": strings.TrimSpace(usage.String())}); err != nil {
		panic(err)
	}
}
