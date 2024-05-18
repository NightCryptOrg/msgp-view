package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tinylib/msgp/msgp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: msgp-view {file}")
		return
	}

	var err error
	err = readFile(os.Args[1])

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

	// Pipe output to jq for pretty-printing
	cmd := exec.Command("jq")
	cmd.Stdout = os.Stdout
	jq, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	cmd.Start()
	_, err = msgp.CopyToJSON(jq, file)
	jq.Close()
	cmd.Wait()
	return err
}
