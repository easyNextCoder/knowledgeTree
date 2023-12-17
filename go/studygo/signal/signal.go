package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func watch() {
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case _, ok := <-stopSignalChan:
			if ok {
				time.Sleep(time.Second * 3)
				fmt.Println("got signal")
				time.Sleep(time.Second * 10)
				fmt.Println("finally done")
			}
		}
	}
}

func main() {
	go watch()
	time.Sleep(time.Second * 10)
}
