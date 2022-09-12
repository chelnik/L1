package main

import (
	"fmt"
)

func main() {
	channnelFirst := make(chan int)
	channnelSecond := make(chan int)

	go func(in chan int) {
		array := [5]int{1, 2, 3, 4, 5}
		for i := 0; i < len(array); i++ {
			in <- array[i]
		}
		// обязательно закрываем канал
		close(in)
	}(channnelFirst)
	go func(out, in chan int) {
		for {
			result, ok := <-out
			if !ok {
				break
			}
			in <- result * result
		}
		// обязательно закрываем канал
		close(in)
	}(channnelFirst, channnelSecond)

	for res := range channnelSecond {
		fmt.Println(res)
	}
}
