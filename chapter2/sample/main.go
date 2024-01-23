package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers" // 短横线_是为了让go对包做初始化操作init
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
