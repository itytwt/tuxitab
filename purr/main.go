package main

import (
	"os"

	"github.com/lubrige/tuxitab/purr/cmd"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rc := cmd.NewRootCmd(cmd.Config{OutputFolder: dir})
	rc.Execute()
}
