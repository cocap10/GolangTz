package main

import (
	"fmt"
)

func fizzbuzz(number int) (msg string) {
	if number%5 == 0 {
		msg = "fizz"
	}
	if number%7 == 0 {
		msg += "buzz"
	}
	return
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d %s\n", i, fizzbuzz(i))
	}
}
