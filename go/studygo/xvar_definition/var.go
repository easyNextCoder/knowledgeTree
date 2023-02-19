package xvar_definition

import "fmt"

func Work() {

	var xint []int
	xint = append(xint, 1) //ok
	fmt.Println(xint)

	var vchan chan int //panic nil chan
	vchan <- 1
	fmt.Println(vchan)

	var xp *[]int
	fmt.Println(xp)
}
