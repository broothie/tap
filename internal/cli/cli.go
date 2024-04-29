package cli

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/broothie/tap"
)

//go:embed help.txt
var help string

type CLI struct {
	*kingpin.Application
	paths    []string
	timeout  time.Duration
	fileMode string
	dirMode  string
}

func New() *CLI {
	cli := &CLI{Application: kingpin.New("tap", strings.TrimSpace(help))}
	cli.Version(tap.Version())
	cli.HelpFlag.Short('h')

	cli.
		Arg("paths", "Paths of files or directories to create.").
		Required().
		StringsVar(&cli.paths)

	cli.
		Flag("timeout", "Command timeout.").
		Default("5s").
		DurationVar(&cli.timeout)

	cli.
		Flag("file-mode", "Mode to use for created files.").
		Short('f').
		Default(fmt.Sprintf("%#o", tap.DefaultFileMode)).
		StringVar(&cli.fileMode)

	cli.
		Flag("dir-mode", "Mode to use for created directories.").
		Short('d').
		Default(fmt.Sprintf("%#o", tap.DefaultDirMode)).
		StringVar(&cli.dirMode)

	return cli
}

func (c *CLI) Paths() []string {
	return c.paths
}

func (c *CLI) Timeout() time.Duration {
	return c.timeout
}

func (c *CLI) FileMode() os.FileMode {
	fileMode, err := strconv.ParseUint(c.fileMode, 8, 64)
	if err != nil {
		panic(err)
	}

	return os.FileMode(fileMode)
}

func (c *CLI) DirMode() os.FileMode {
	dirMode, err := strconv.ParseUint(c.dirMode, 8, 64)
	if err != nil {
		panic(err)
	}

	return os.FileMode(dirMode)
}
