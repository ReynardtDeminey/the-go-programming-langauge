package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("Go"))
	c2 := sha256.Sum256([]byte("Away"))
	c3 := fmt.Sprintf("%x", c1)
	c4 := fmt.Sprintf("%x", c2)
	fmt.Println(len(c3))
	fmt.Println(len(c4))
	r := diff(c3, c4)
	fmt.Println(r)
}

func diff(s1, s2 string) int {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}
