package main

import (
	"fmt"
	"os"
)

const Version = "1"

func handleVersion() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "version") {
		fmt.Println("cleanly v" + Version)
		os.Exit(0)
	}
}