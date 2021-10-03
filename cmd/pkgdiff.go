package main

import (
	"fmt"
	"github.com/cbuschka/pkgdiff/internal"
	"os"
)

func main() {

	err := internal.Run()
	if err != nil {
		fmt.Printf("Failed to run package diff: %v", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
