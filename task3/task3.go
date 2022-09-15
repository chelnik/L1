package main

import (
	"fmt"
)

func main() {
	//t := time.Now()
	var sum int
	message := make(chan int)
	array := [5]int{2, 4, 6, 8, 10}
	go func() {
		for _, number := range array {
			message <- number * number
			//time.Sleep(time.Millisecond * 500)
		}
		close(message)
	}()

	for msg := range message {
		sum += msg
	}
	fmt.Println(sum)
	//fmt.Println(time.Since(t).Nanoseconds())
}
