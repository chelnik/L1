package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGQUIT:
				fmt.Println("Caught:", sig)
				return
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Millisecond)
	}
}
