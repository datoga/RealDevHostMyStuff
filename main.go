package main

import (
	"os"
)

var debug bool

func main() {

	if *flagP == "" {
		panic("Route not provided")
	}

	dirOrFile := *flagP

	debug = *flagD

	if _, err := os.Stat(dirOrFile); os.IsNotExist(err) {
		panic("Route " + dirOrFile + " does not exist")
	}

	uploadToS3(dirOrFile)
}
