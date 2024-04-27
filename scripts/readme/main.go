package main

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/broothie/tap"
)

//go:embed README.md.sprintf
var readmeFormat string

func main() {
	output, err := exec.Command("tap", "-h").CombinedOutput()
	if err != nil {
		panic(err)
	}

	if err := tap.Tap("README.md", tap.Content([]byte(fmt.Sprintf(readmeFormat, output)))); err != nil {
		panic(err)
	}
}
