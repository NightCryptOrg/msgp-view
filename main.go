package main

import (
	"fmt"
	"os"

	"github.com/tinylib/msgp/msgp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: mgsp-view {file}")
		return
	}

	var err error
	err = readFile(os.Args[1])
	fmt.Println()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = msgp.CopyToJSON(os.Stdout, file)
	return err
}
