package main

import (
	"os"

	"github.com/fadeev/pluto/cmd/plutod/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
