package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// собственно воркер
func startWorker(ctx context.Context, wg *sync.WaitGroup, workerNum int, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case data := <-in:
			fmt.Printf("%d %d\n", workerNum, data)
		case <-ctx.Done():
			fmt.Println("close worker", workerNum)
			return
		}
	}
}

func scannerFunc() int {
	scanner := bufio.NewScanner(os.Stdout)
	scanner.Scan()
	number, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Println(err)
	}
	return int(number)
}

func main() {
	goroutinesNum := scannerFunc()
	// создаем контекст
	ctx, cancel := context.WithCancel(context.Background())
	//создаем канал для получения сигнала
	sigs := make(chan os.Signal, 1)
	// подписываемся на sigint
	signal.Notify(sigs, syscall.SIGINT)
	// создаем канал для чтобы передавать данные в воркер
	workerInput := make(chan int, goroutinesNum)
	//создаем WaitGroup чтобы корректно обработать закрытие воркеров
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(ctx, wg, i, workerInput)
	}
	i := 0

	for {
		select {
		case sig := <-sigs:
			// просто выводим сам сигнал закрываем канал и ждем закрытия всех воркеров
			fmt.Println()
			fmt.Println(sig)
			cancel()
			//close(workerInput)
			wg.Wait()
			return
		default:
			workerInput <- i
			i++
		}
		// делаем специально задержку чтобы видеть корректно данные
		time.Sleep(time.Millisecond * 100)
	}
}
