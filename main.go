package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	timeout := flag.Duration("timeout", 5*time.Second, "Command timeout")
	fileModeFlag := flag.String("mode", "0666", "File mode")
	dirModeFlag := flag.String("mode", "0777", "Directory mode")
	flag.Parse()

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

	if err := run(ctx, flag.Args(), os.Stdin, os.FileMode(fileMode), os.FileMode(dirMode)); err != nil {
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
			if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
				return fmt.Errorf("creating directory: %w", err)
			}

			if strings.HasSuffix(path, string(os.PathSeparator)) {
				return nil
			}

			if err := os.WriteFile(path, content.Bytes(), fileMode); err != nil {
				return fmt.Errorf("creating directory: %w", err)
			}

			return nil
		})
	}

	return group.Wait()
}
