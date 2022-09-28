package main

import (
	"fmt"
)

func main() {
	var sum int
	message := make(chan int)
	array := [5]int{2, 4, 6, 8, 10}
	go func() {
		for _, number := range array {
			message <- number * number
		}
		close(message)
	}()

	for msg := range message {
		sum += msg
	}
	fmt.Println(sum)
}
