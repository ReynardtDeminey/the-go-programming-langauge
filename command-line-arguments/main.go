package main

import (
	"fmt"
	"os"
	"strings"
)

// Implementation with a normal for loop
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	altArgs()
	joinArgs()

	// Just using the Println function
	fmt.Println(os.Args[1:])
}

// Implementation by ranging over the command-line arguments
func altArgs() {
	var s, sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
}

// Implementation by using the Join function from the strings package
func joinArgs() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
