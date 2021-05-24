package main

import (
	"os"

	"github.com/dedi27/compellentbeat/cmd"

	_ "github.com/dedi27/compellentbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
