package main

import "sync"

func main() {
	var wg sync.WaitGroup
	chan_first := make(chan int)
	chan_second := make(chan int)
	chan_third := make(chan int)
	for i := 0; i < 10; i++ {
		go func(wg &sync.WaitGroup) {

		}(*wg, i int)
	}


	go func(wg &sync.WaitGroup) {

	}(*wg)
}
