package main

import (
	"fmt"
	"os"

	"github.com/djamboe/mtools-location-service/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
