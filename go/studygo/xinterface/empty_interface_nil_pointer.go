package xinterface

import "fmt"

type EmptyInterfaceNilPointer struct {
}

func emptyInterfaceNilPointer() {
	var x interface{}
	var y *EmptyInterfaceNilPointer

	fmt.Println(x == y)
}
