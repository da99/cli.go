
package main

import (
	"fmt"
	"github.com/da99/cli.go/args"
	"os"
)

func main() {
	fmt.Printf("ARGS: %v\n", os.Args)
	if args.IsMatch("my files", 0) {
		fmt.Printf("CAPTURE: %v\n", args.CAPTURE)
		return
	}

	if args.IsMatch("my files", 2) {
		fmt.Printf("CAPTURE: %v\n", args.CAPTURE)
		return
	}

	fmt.Printf("CAPTURE: %v\n", args.CAPTURE)
	fmt.Printf("Unknown args: %v\n", os.Args[1:])
}
