package main

import (
	"fmt"

	"github.com/ReynardtDeminey/the-go-programming-langauge/hamming/hamming"
)

func main() {
	hd := hamming.Hamming("AAAAA", "BBBBB")
	fmt.Println(hd)
	hd2 := hamming.Hamming("ABCDE", "ABCDD")
	fmt.Println(hd2)
}
