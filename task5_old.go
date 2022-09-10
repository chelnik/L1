package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// берем число
	channel := make(chan int)
	scanner := bufio.NewScanner(os.Stdout)
	scanner.Scan()
	nSeconds, _ := strconv.ParseFloat(scanner.Text(), 64)
	t := time.Now()
	go func(channel chan int) {
		fmt.Print(<-channel, "\n")
	}(channel)
	for i := 0; time.Since(t).Seconds() < nSeconds; i++ {
		channel <- i
	}
	context.WithTimeout()
}
