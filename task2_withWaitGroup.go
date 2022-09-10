package main

import (
	"fmt"
	"sync"
)

func main() {
	//t := time.Now()
	array := [5]int{2, 4, 6, 8, 10}
	//создаем буферизировнный канал
	jekiChan := make(chan int, 5)
	wg := &sync.WaitGroup{}
	for _, number := range array {
		wg.Add(1)
		go func(squadChan chan<- int, number int, wg *sync.WaitGroup) {
			defer wg.Done()
			squadChan <- number * number
		}(jekiChan, number, wg)
	}
	wg.Wait()
	close(jekiChan)
	for i := range jekiChan {
		fmt.Println(i)
	}
	//fmt.Println(time.Since(t).Nanoseconds())
}
