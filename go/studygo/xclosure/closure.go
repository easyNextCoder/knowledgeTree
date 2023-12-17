package main

import (
	"fmt"
	"time"
)

func parent() func() {
	i := 10 //局部变量会提到heap中，成为内外部沟通的桥梁
	return func() {
		i++
		fmt.Println(i)
	}
}

func main() {
	son := parent()
	allDoneCh := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("son1 print")
			son()
		}
		allDoneCh <- 1
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("son2 print")
			son()
		}
		allDoneCh <- 1
	}()

	go func() {
		var count int = 0
		for {
			select {
			case <-allDoneCh:
				count++
				if count == 2 {
					go func() {
						newSon := parent()
						for i := 0; i < 5; i++ {
							fmt.Println("newSon print")
							newSon()
						}
					}()
					return
				}
			}
		}
	}()

	time.Sleep(time.Second * 2)
}
