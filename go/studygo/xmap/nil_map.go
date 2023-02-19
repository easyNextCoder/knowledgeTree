package xmap

import "fmt"

func nilMapWork() {
	var x map[int]int
	fmt.Println("nilMapWork", x[1])
	x[1] = 1

}

func assMapWork() {
	x := map[int]int{}
	fmt.Println("nilMapWork", x[1])
	x[1] = 1
}
