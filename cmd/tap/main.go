package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/broothie/tap"
	"github.com/broothie/tap/internal/cli"
	"golang.org/x/sync/errgroup"
)

func main() {
	cli := cli.New()
	if _, err := cli.Parse(os.Args[1:]); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, cli.Timeout())
	defer cancel()

	if err := run(ctx, cli.Paths(), os.Stdin, cli.FileMode(), cli.DirMode()); err != nil {
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
