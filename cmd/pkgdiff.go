package main

import (
	"fmt"
	"github.com/cbuschka/go-pkgdiff/internal/command"
	"os"
)

func main() {

	err := command.Run()
	if err != nil {
		fmt.Printf("Failed to run package diff: %v", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
