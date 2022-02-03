# Touchd - A friendlier version of touch

## Background

I often find myself wanting to create a file inside of a non existing directory.
To accomplish this I usually type something like `mkdir foo && touch foo/bar`, but this is far too many keystrokes for my liking!

## Usage

`touchd` solves this problem by combining `mkdir -p` and `touch` into a single command. From `touchd -help`:

```
Usage: touchd [OPTION] FILE...

Update the access and modification times of each FILE to the current time.

A FILE argument that does not exist is created empty
If any FILE argument contains parent directories that do not exist, they are created automatically.

Options:
-help
        Print this help message and quit
```

## Installing

Install with the `go install` tool.

```
$ go install github.com/penguingovernor/cmd/touchd
```