# `tap`

## Installation

```
go install github.com/broothie/tap/cmd/tap@latest
```

## Usage

```
usage: tap [<flags>] <paths>...

Given a path, tap creates the file or directory at the given path, along with
any nonexistent directories in the path.

Example usage:

  $ tap file.txt                        # Creates 'file.txt'.
  $ tap path/to/file.txt                # Creates dirs 'path/', 'path/to/', and file 'path/to/file.txt'.
  $ tap path/to/dir/                    # Creates dirs 'path/', 'path/to/', and 'path/to/dir/'.
  $ echo "Hello, World!" | tap file.txt # Creates 'file.txt' and writes stdin to it.


Flags:
  -h, --[no-]help         Show context-sensitive help (also try --help-long and
                          --help-man).
      --[no-]version      Show application version.
      --timeout=5s        Command timeout.
  -f, --file-mode="0644"  Mode to use for created files.
  -d, --dir-mode="0755"   Mode to use for created directories.

Args:
  <paths>  Paths of files or directories to create.
```
