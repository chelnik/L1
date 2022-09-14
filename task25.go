package main

import (
	"fmt"
	"time"
)

//func mySleep(duration time.Duration) {
//	// фунцкия возвращает канал
//	var timer = time.After(duration)
//	<-timer
//}
//
//func main() {
//	//time.Sleep(-1)
//	mySleep(10000000000 / 5)
//	fmt.Println("go")
//}

func mySleep(duration time.Duration) {
	// фунцкия возвращает канал
	ticker := time.NewTicker(duration)
	_ = <-ticker.C
}

func main() {
	//time.Sleep(-1)
	mySleep(10000000000 / 5)
	fmt.Println("go")
}
