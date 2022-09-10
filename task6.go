package main

import (
	"fmt"
	"time"
)

func main() {
	exitChannel := make(chan bool)
	go func(in <-chan bool) {
		for {
			select {
			case <-in:
				fmt.Println("\ngoodbye")
				return
			default:
				fmt.Print(".")
				time.Sleep(time.Millisecond * 5)
			}
		}

	}(exitChannel)
	time.Sleep(time.Millisecond * 1000)
	exitChannel <- true
	close(exitChannel)
}
