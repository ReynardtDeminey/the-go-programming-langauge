// Package main provides utilities for printing the given command-line arguments to the screen
package main

import (
	"os"
	"strings"
)

func main() {
	printArgs()
	printArgsRange()
	printArgsJoin()
}

// printArgs prints the command-line arguments to the screen using a for loop
func printArgs() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func printArgsRange() string {
	var s, sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	return s
}

func printArgsJoin() string {
	s := strings.Join(os.Args[1:], " ")
	return s
}
