package main

import (
	"os"

	"github.com/avarian/go-template/cli/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
