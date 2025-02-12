package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// need to import the module with `go mod tidy` before running with `go run .`
