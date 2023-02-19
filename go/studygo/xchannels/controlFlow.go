package xchannels

import "fmt"

func testFor() {
	fmt.Println("test controlFlow")
	for i := 0; i < 10; i++ {
		fmt.Println("for", i)
	}
}
