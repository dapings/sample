package main

import (
	"log"
	"os"
	
	_ "github.com/dapings/sample/matchers"
	"github.com/dapings/sample/search"
)

var logPrefix = "dapings sample "

func init() {
	log.SetPrefix(logPrefix)
	// change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	// perform the search for the specified term.
	search.Run("president")
}
