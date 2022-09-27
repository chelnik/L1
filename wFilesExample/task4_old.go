package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}
func creater(ch chan int, num int) {
	ch <- num
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	t := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("need #jobs to #workers!\nВведите количество воркеров")
		os.Exit(1)
	}
	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	channel := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		for i := 0; i < nWorkers; i++ {
			wg.Add(1)
			go creater(channel, i)
			go worker(channel, wg)
		}
	}()
	done := <-sigs
	wg.Wait()
	fmt.Println("Got signal", done)
	fmt.Println(time.Since(t).Nanoseconds())
}

// не могу решить проблему дедлока с помощью WaitGroup
