package main

import (
	"fmt"
	"os"

	"github.com/abh2020-B2/todo-cli/internal/cli"
)

func main() {
	// Run the CLI application.
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
