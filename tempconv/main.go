package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ReynardtDeminey/the-go-programming-langauge/tempconv/tempconv"
)

func main() {
	for _, v := range os.Args[1:] {
		t, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v Degrees Celsius:\n", t)
		fmt.Println(tempconv.CtoF(tempconv.Celsius(t)))

		fmt.Printf("%v Degrees Fahrenheit:\n", t)
		fmt.Println(tempconv.FtoC(tempconv.Fahrenheit(t)))
	}
}
