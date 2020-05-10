package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	altOs()

}

func altOs() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%v\n", os.Args[i])
	}
}
