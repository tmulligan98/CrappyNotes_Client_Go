package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("help", false, "Display help message.")
	version := flag.Bool("version", false, "Display version message.")

	flag.Parse()

	if *help {
		fmt.Println("Blah blah help blah blah")
		os.Exit(0)
	}

	if *version {
		fmt.Println("Blah blah version blah blah")
		os.Exit(0)
	}

	if !*help && !*version {
		fmt.Println("Blah blah default message blah blah")
		os.Exit(0)
	}

}
