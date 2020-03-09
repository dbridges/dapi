package main

import (
	"fmt"
	"os"
)

// Version is set at build time
var Version string

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Println("dapi", Version)
			return
		} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println("dapi By Dan Bridges")
			fmt.Println("See https://github.com/dbridges/dapi")
			return
		}
	}
}
