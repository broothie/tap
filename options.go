package tap

import (
	"os"

	"github.com/broothie/option"
)

const (
	DefaultFileMode os.FileMode = 0644
	DefaultDirMode  os.FileMode = 0755
)

type Options struct {
	content  []byte
	fileMode os.FileMode
	dirMode  os.FileMode
}

func defaultOptions() option.Options[Options] {
	return option.Options[Options]{
		FileMode(DefaultFileMode),
		DirMode(DefaultDirMode),
	}
}

func Content(content []byte) option.Func[Options] {
	return func(options Options) (Options, error) {
		options.content = content
		return options, nil
	}
}

func FileMode(fileMode os.FileMode) option.Func[Options] {
	return func(options Options) (Options, error) {
		options.fileMode = fileMode
		return options, nil
	}
}

func DirMode(dirMode os.FileMode) option.Func[Options] {
	return func(options Options) (Options, error) {
		options.dirMode = dirMode
		return options, nil
	}
}
