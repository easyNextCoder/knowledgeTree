package xpanic

import "fmt"

func directPanic() {
	panic("directPanic")
}

func panicRecover() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicRecover err(%s)", e)
		}
	}()
	work()
}

func work() {
	panic("panicRecover")
}
