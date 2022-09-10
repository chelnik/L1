package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	waitChannel := make(chan bool)
	stx, _ := context.WithTimeout(context.Background(), time.Second)
	go func(stx context.Context, waitChannel chan bool) {
		for {
			select {
			case <-stx.Done():
				fmt.Println("\ngoodbye")
				waitChannel <- true
				return
			default:
				fmt.Print(".")
				time.Sleep(time.Millisecond * 5)
			}
		}
	}(stx, waitChannel)
	<-waitChannel
}
