package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/penguingovernor/touchd/pkg/touchd"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] FILE...\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "Update the access and modification times of each FILE to the current time.\n")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "A FILE argument that does not exist is created empty\n")
		fmt.Fprintf(os.Stderr, "If any FILE argument contains parent directories that do not exist, they are created automatically.\n")
		fmt.Fprintln(os.Stderr)
		fmt.Println("Options:")
		fmt.Fprintf(os.Stderr, "-help\n\tPrint this help message and quit\n")
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "%s: missing file operand", os.Args[0])
		os.Exit(1)
	}
	if err := touchd.CreateFiles(flag.Args()...); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
