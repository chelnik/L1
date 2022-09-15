package main

import (
	"fmt"
)

func main() {
	//t := time.Now()
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
		fmt.Println(msg)
	}
	//for {
	//	msg, open := <-message
	//	if !open {
	//		break
	//	}
	//	fmt.Println(msg)
	//}
	//fmt.Println(time.Since(t).Nanoseconds())
}
