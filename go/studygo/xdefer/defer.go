package main

import (
	"fmt"
	"time"
)

func main() {

	defer func() {
		fmt.Println("end of main")
	}()

	defer func() {
		fmt.Println("pre end of main")
	}()

	go func() {
		var i int
		defer func() {
			fmt.Println(i)
		}()
		i = 10
		fmt.Println("end of go func")
	}()
	time.Sleep(100 * time.Microsecond)
}
