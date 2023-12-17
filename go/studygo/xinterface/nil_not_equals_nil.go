package xinterface

import "fmt"

type EmptyInterfaceNilPointer struct {
}

func emptyInterfaceNilPointer() {
	var x interface{}
	var y *EmptyInterfaceNilPointer
	var z EmptyInterfaceNilPointer

	//x 是 nil
	//y 是 nil
	//但是两个nil是不相等的

	fmt.Println(x == y, x, y, z)
}
