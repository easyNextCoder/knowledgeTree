package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectNew() {
	// one way is to have a value of the type you want already
	a := 1
	// reflect.New works kind of like the built-in function new
	// We'll get a reflected pointer to a new int value
	intPtr := reflect.New(reflect.TypeOf(a))
	// Just to prove it
	b := intPtr.Elem().Interface().(int)
	// Prints 0
	fmt.Println(b)

	// We can also use reflect.New without having a value of the type
	var nilInt *int
	intType := reflect.TypeOf(nilInt).Elem()
	intPtr2 := reflect.New(intType)
	// Same as above
	c := intPtr2.Elem().Interface().(int)
	// Prints 0 again
	fmt.Println(c)
	fmt.Println(intPtr2)
}

type A struct {
	AA int
	AB string
}

func reflectNew1(a *A) {
	inputData := CommandWrapper{
		Data: &A{
			AA: 9999,
			AB: "xuyongkang",
		},
		Result:  999,
		Message: "目前运行正常！",
	}

	v, _ := json.Marshal(inputData)
	fmt.Println("after marshal: ", string(v))
	aType := reflect.TypeOf(a).Elem()
	pa := reflect.New(aType).Interface()

	wrapperPt := new(CommandWrapper)
	wrapperPt.Data = pa

	json.Unmarshal(v, wrapperPt)
	fmt.Println(inputData, pa, wrapperPt)
}
