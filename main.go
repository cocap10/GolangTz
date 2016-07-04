package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		var message string
		if i%5 == 0 {
			message = "fizz"
		}
		if i%7 == 0 {
			message += "buzz"
		}
		fmt.Printf("%d %s\n", i, message)
	}
}
