# `tap`

## Installation

```
go install github.com/broothie/tap/cmd/tap@latest
```

## Usage

```
usage: tap [<flags>] <paths>...

Creates files and directories.

Example usage:

  $ tap file.txt                        # Creates 'file.txt'.
  $ tap dir/file.txt                    # Creates 'file.txt', and 'dir/' if it doesn't already exist.
  $ tap path/to/file.txt                # Creates 'path/', 'path/to/', and 'path/to/file.txt'.
  $ tap path/to/dir/                    # Creates 'path/', 'path/to/', and 'path/to/dir/'.
  $ echo "Hello, World!" | tap file.txt # Creates 'file.txt' and writes stdin to it.


Flags:
  -h, --[no-]help         Show context-sensitive help (also try --help-long and
                          --help-man).
  -V, --[no-]version      Show application version.
      --timeout=5s        Command timeout.
  -f, --file-mode="0644"  Mode to use for created files.
  -d, --dir-mode="0755"   Mode to use for created directories.

Args:
  <paths>  Paths of files or directories to create.
```
