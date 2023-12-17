package main

import "fmt"

type P struct {
	name string
}

func (self P) print() { //如果没有(self P)那结构体P变量调用的是(self *P)
	fmt.Println(self.name)
	self.name = "after print"
}

func (self *P) PlayCat() {
	fmt.Println("so pain")
}

func main() {
	var p P
	p.print()
	p.print()
	p.PlayCat()
	fmt.Println("final print", p.name)

	fmt.Println("函变的默", fn)
}

var fn func(int)

