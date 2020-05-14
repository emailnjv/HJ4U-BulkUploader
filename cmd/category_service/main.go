package main

import (
	"fmt"
	"os"
)

func init() {
}

func main() {
	exit(run())
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}
	os.Exit(1)
}

func run() error {
}

func printUsage() {
	fmt.Println("")
	os.Exit(0)
}