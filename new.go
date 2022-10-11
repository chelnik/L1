package main

import (
	"fmt"
	"sync"
	"time"
)
import "os"
import "syscall"
import "os/signal"

func worker(worker_num int, ch <-chan int, channelForSecond chan int, wg *sync.WaitGroup) {
	for number := range ch {
		time.Sleep(time.Second)
		fmt.Print("Worker  ", worker_num, "\n", "Received: ", number, "\n")
		channelForSecond <- number
	}

	fmt.Printf("worker %d done ", worker_num)
	wg.Done()
}

func workerSecond(worker_num int, ch <-chan int, wg *sync.WaitGroup) {
	//ch := make(chan bool)

	for number := range ch {
		time.Sleep(time.Second)
		fmt.Print("Worker  ", worker_num, "\n", "Received: ", number, "\n")
	}
	fmt.Printf("worker %d done ", worker_num)

	wg.Done()
}

func main() {
	n := 2
	var wg sync.WaitGroup
	wg.Add(n)
	ch := make(chan int)
	channelForSecond := make(chan int)
	go worker(1, ch, channelForSecond, &wg)
	go workerSecond(2, channelForSecond, &wg)

	some_value := 0
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT)
	for {
		select {
		case <-sign:
			fmt.Printf("worker %d done ", 1)

			close(ch)
			return
		default:
			some_value++
			ch <- some_value
		}
	}
	wg.Wait()
}
