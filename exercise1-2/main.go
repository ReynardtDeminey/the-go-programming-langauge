// Package main provides utilities to print the given command-line arguments to the screen
package main

import (
	"fmt"
	"os"
)

func main() {
	printArgs()
}

// printArgs prints the given command-line arguments and their index to the screen
func printArgs() {
	s := os.Args[1:]
	for i, v := range s {
		fmt.Println(i, v)
	}

}
