package xvar_definition

import "fmt"

func vart() {
	var p []int

	b := []int{}

	fmt.Println(p, b)
}

func maket() {
	p := new([]int)

	b := []int{}

	fmt.Println(p, b)
}

func newt() {
	p := make([]int, 0)

	b := []int{}

	fmt.Println(p, b)
}
