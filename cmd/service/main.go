package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wilorios/microservice-template-go2/internal/application"
)

const (
	// exitFail exit code to use when the program fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	log.Println("the application has finished")
}

func run(args []string, stdouot io.Writer) error {
	newApplication := application.New(args)
	if err := newApplication.Start(); err != nil {
		log.Println("fatal", err)
		return err
	}
	return nil
}
