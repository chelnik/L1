package main

import (
	"fmt"
	"sync"
)

// using mutex
func concurrencyWriteInMap() (myMap map[int]string) {
	mutex := sync.Mutex{}
	wg := &sync.WaitGroup{}
	myMap = make(map[int]string)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			mutex.Lock()
			myMap[i] = fmt.Sprintf("new string type %T value %d", i, i)
			mutex.Unlock()
		}(wg, i)
	}
	wg.Wait()
	return
}

func main() {
	myMap := concurrencyWriteInMap()
	for i := 0; i < len(myMap); i++ {
		fmt.Println(myMap[i])
	}
}
