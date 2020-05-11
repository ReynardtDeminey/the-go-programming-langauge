// Package main provides utilities to print the given command-line arguments to the screen
package main

import (
	"fmt"
	"os"
)

func main() {
	printArgs()
	printCmd()
}

// printArgs prints the given command-line arguments to the screen
func printArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

// printCmd prints the arguments and the command to the screen
func printCmd() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Printf("Command: %v, result: %v\n", os.Args[0], s)
}
