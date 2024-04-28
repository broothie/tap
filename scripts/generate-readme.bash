#! /usr/bin/env bash

cat <<- "README" > README.md
# `tap`

## Installation

```
go install github.com/broothie/tap/cmd/tap@latest
```

## Usage

```
README

tap -h 2>> README.md

echo '```' >> README.md
