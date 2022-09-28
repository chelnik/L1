package main

import (
	"fmt"
)

func main() {
	message := make(chan int)
	array := [5]int{2, 4, 6, 8, 10}
	go func() {
		for _, number := range array {
			message <- number * number
		}
		close(message)
	}()
	for msg := range message {
		fmt.Println(msg)
	}
}

// более многословный вариант чтения из канала
//for {
//	msg, open := <-message
//	if !open {
//		break
//	}
//	fmt.Println(msg)
//}
