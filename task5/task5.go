package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
func main() {
	// берем число
	channel := make(chan int, 1)
	//channelExit := make(chan bool)
	scanner := bufio.NewScanner(os.Stdout)
	scanner.Scan()
	nSeconds, _ := strconv.ParseFloat(scanner.Text(), 64)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(nSeconds)*time.Second)

	go func(channel chan int) {
		for {
			fmt.Print(<-channel, "\n")
		}

	}(channel)

	ticker := time.NewTicker(100 * time.Millisecond)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("bye")
			break LOOP
		default:
			t := <-ticker.C
			channel <- t.Second()
		}
	}
}

//go func(ctx context.Context, channel chan int) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("bye")
//			return
//		default:
//			fmt.Print(<-channel, "\n")
//		}
//	}
//
//}(ctx, channel)
