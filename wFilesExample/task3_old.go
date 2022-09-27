package main

import (
	"fmt"
)

func main() {
	//t := time.Now()
	var sum int
	array := [5]int{2, 4, 6, 8, 10}
	message := make(chan int, len(array))
	go func() {
		for _, number := range array {
			message <- number * number
		}
		close(message)
	}()
	for {
		msg, ok := <-message
		if !ok {
			break
		}
		sum += msg
	}
	// for msg := range message {
	// 	sum += msg
	// }

	fmt.Println(sum)
	//fmt.Println(time.Since(t).Nanoseconds())
}

//i := 0; i < len(array); i++
