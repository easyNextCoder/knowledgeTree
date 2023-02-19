package arr

import "fmt"

type ArrFuncWork struct {
	v [10]func()
}

func arrFuncWork() {
	var x ArrFuncWork
	fmt.Println(x)
}
