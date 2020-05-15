// Coinflip displays the ratio of heads vs tails
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	times := os.Args[1]
	timesInt, err := strconv.ParseInt(times, 0, 32)
	if err != nil {
		panic(err)
	}
	res := generateRatio(int(timesInt))
	fmt.Println(res)
}

// generateNumber generates a random number between 0 and 1
func generateNumber() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(2)
	return num
}

// generateRatio determines the ratio of both heads and tails out of a given number of times
func generateRatio(times int) string {
	var heads int
	var tails int
	for i := 1; i <= times; i++ {
		num := generateNumber()
		if num == 1 {
			heads++
		} else if num == 0 {
			tails++
		}
	}
	return fmt.Sprintf("Out of %v times, heads was %v and tails %v", times, heads, tails)
}
