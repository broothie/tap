package tap

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/broothie/option"
)

//go:embed VERSION
var version string

func Version() string {
	return strings.TrimSpace(version)
}

func Tap(path string, options ...option.Option[Options]) error {
	if strings.HasSuffix(path, string(os.PathSeparator)) {
		return Dir(path, options...)
	}

	return File(path, options...)
}

func Dir(path string, options ...option.Option[Options]) error {
	opts, err := append(defaultOptions(), options...).Apply(Options{})
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path, opts.dirMode); err != nil {
		return fmt.Errorf("creating directories %q: %w", path, err)
	}

	return nil
}

func File(path string, options ...option.Option[Options]) error {
	opts, err := append(defaultOptions(), options...).Apply(Options{})
	if err != nil {
		return err
	}

	if err := Dir(filepath.Dir(path), options...); err != nil {
		return err
	}

	if err := os.WriteFile(path, opts.content, opts.fileMode); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
