package main

import (
	"fmt"
	"sync"
)

// make mutex with chan
func concurrencyWriteInMap() (myMap map[int]string) {
	mutexChan := make(chan struct{}, 1)
	wg := &sync.WaitGroup{}
	myMap = make(map[int]string)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			mutexChan <- struct{}{}
			myMap[i] = fmt.Sprintf("new string type %T value %d", i, i)
			<-mutexChan
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
