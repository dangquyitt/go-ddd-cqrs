package main

import (
	"fmt"
	"os"

	"github.com/dangquyitt/go-ddd-cqrs/cmd/myapp/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Some error occurred during execute app. Error: %v\n", err)

		os.Exit(2)
	}
}
