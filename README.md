# goharm
Rebased Harmonogram CLI (written in Golang)

## Usage

```bash
$ build/goharm-darwin-amd64
Rebased Harmonogram CLI
Usage of build/goharm-darwin-amd64:
  -v	Show version number

	goharm [SUBCOMMAND]

expected 'login', 'logout', 'config' or 'time-logs' subcommands
```

Each subcommand has it's own `-help` parameter.

## Development

### How to run tests?

```bash
make test
```

### How to release?

Use command with make with appropriate TAG for version:

```bash
TAG=0.0.X make clean test tag build release
```

(c) Rafał "RaVbaker" Piekarski 2020