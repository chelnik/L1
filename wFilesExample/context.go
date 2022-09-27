package main

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)
	time.Sleep(time.Millisecond * 5000)
}

func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}
