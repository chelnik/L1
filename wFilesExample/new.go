package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const goroutinesNum = 100000

func startWorker(ctx context.Context, workerNum int, in <-chan int) {
	for {
		select {
		case data := <-in:
			fmt.Printf("%d %d\n", workerNum, data)
		case <-ctx.Done():
			return
		}
	}
	//printFinishWork(workerNum)
}

//func printFinishWork(number int) {
//	println(number)
//}
func main() {
	ctx, finish := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	worketInput := make(chan int, 2) // попробуйте увеличить размер канала

	for i := 0; i < goroutinesNum; i++ {
		go startWorker(ctx, i, worketInput)
	}
	i := 0

	for {
		select {
		case sig := <-sigs:
			fmt.Println()
			fmt.Println(sig)
			finish()
			return
		default:
			worketInput <- i
			i++
		}
		time.Sleep(time.Millisecond * 10)
	}
	//close(worketInput) // попробуйте закомментировать
	// time.Sleep(time.Millisecond)
}
