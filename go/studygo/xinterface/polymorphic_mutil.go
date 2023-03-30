package xinterface

import "fmt"

type Swimmer interface {
	swimming()
}

type Flyer interface {
	fly()
}

type AngryBird struct {
	name string
}

func (a AngryBird) swimming() {
	fmt.Println(a.name, "swimming...")
}

func (a AngryBird) fly() {
	fmt.Println(a.name, "flying...")
}

func swimmerAccepter(s Swimmer) {
	if sFlyer, ok := s.(Flyer); ok {
		sFlyer.fly()
	}
}

type LoveBird struct {
	Swimmer
	Flyer
	tag string
}

func playGround() {

	//多态
	var angryBird1 AngryBird = AngryBird{name: "red"}
	var i interface{} = angryBird1

	if flyer, ok := i.(Flyer); ok {
		flyer.fly()
	}

	if swimmer, ok := i.(Swimmer); ok {
		swimmer.swimming()
	}

	//赋值
	var swimmer Swimmer = AngryBird{name: "blue"}

	swimmerAccepter(swimmer)

	var angryBird AngryBird = AngryBird{name: "angryBird"}

	swimmerAccepter(angryBird) //经过函数参数的时候，会发生类型转换，但是携带的类型信息和变量不会丢失

	//panic
	var loveBird LoveBird

	swimmerAccepter(loveBird) //含有这个类型但是调用没有实现的方法会panic

}
