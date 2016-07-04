package main

import (
	"fmt"
)

func fizzbuzz(numbers chan int, messages chan string) {
	for {
		number := <-numbers
		var msg string
		if number%5 == 0 {
			msg = "fizz"
		}
		if number%7 == 0 {
			msg += "buzz"
		}
		messages <- msg
	}
}

func main() {
	numbers := make(chan int)
	messages := make(chan string)
	go fizzbuzz(numbers, messages)
	for i := 0; i < 100; i++ {
		numbers <- i
		msg := <-messages
		fmt.Printf("%d %s\n", i, msg)
	}
}
