package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/alecthomas/kingpin/v2"
	"github.com/broothie/tap"
	"golang.org/x/sync/errgroup"
)

//go:embed help.txt
var help string

func main() {
	cli := kingpin.New("tap", help)
	cli.Version(tap.Version())
	cli.HelpFlag.Short('h')

	paths := cli.
		Arg("paths", "Paths of files or directories to create.").
		Required().
		Strings()

	timeout := cli.
		Flag("timeout", "Command timeout.").
		Default("5s").
		Duration()

	fileModeFlag := cli.
		Flag("file-mode", "Mode to use for created files.").
		Short('f').
		Default(strconv.Itoa(int(tap.DefaultFileMode))).
		String()

	dirModeFlag := cli.
		Flag("dir-mode", "Mode to use for created directories.").
		Short('d').
		Default(strconv.Itoa(int(tap.DefaultDirMode))).
		String()

	if _, err := cli.Parse(os.Args[1:]); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fileMode, err := strconv.ParseUint(*fileModeFlag, 8, 32)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	dirMode, err := strconv.ParseUint(*dirModeFlag, 8, 32)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, *timeout)
	defer cancel()

	if err := run(ctx, *paths, os.Stdin, os.FileMode(fileMode), os.FileMode(dirMode)); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, input *os.File, fileMode, dirMode os.FileMode) error {
	info, err := input.Stat()
	if err != nil {
		return fmt.Errorf("statting input: %w", err)
	}

	var content bytes.Buffer
	if info.Size() > 0 {
		if _, err := input.WriteTo(&content); err != nil {
			return fmt.Errorf("writing to file: %w", err)
		}
	}

	group, ctx := errgroup.WithContext(ctx)
	for _, path := range args {
		if ctx.Err() != nil {
			break
		}

		group.Go(func() error {
			if err := tap.Tap(
				path,
				tap.Content(content.Bytes()),
				tap.FileMode(fileMode),
				tap.DirMode(dirMode),
			); err != nil {
				return err
			}

			fmt.Println("tap", path)
			return nil
		})
	}

	return group.Wait()
}
