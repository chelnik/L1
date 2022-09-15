package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

// Реализация через канал вместо мьютекса
type Counter struct {
	count int
}

func chanAsMutex() {
	var counter Counter
	//создаем буферизированный канал, чтобы положить в него пустую структуру
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 5)
			// кладем пустую структуру
			mutexChan <- struct{}{}

			counter.count++
			fmt.Println(counter)
			// очищаем буф канал
			<-mutexChan
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func main() {
	t := time.Now()
	chanAsMutex()
	fmt.Println(time.Since(t).Milliseconds())
}
