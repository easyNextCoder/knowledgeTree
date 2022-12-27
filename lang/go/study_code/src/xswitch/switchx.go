package xswitch

import "fmt"

func invokeSwitchx() {
	res := switchx()
	fmt.Println(res)
}

//switch中的每个case自带一个break
func switchx() int {
	val := 0
	var x int
	switch val {
	case 1:
		x = 101
		fmt.Println("work")
	case 2:
		x = 102
		fmt.Println(2)
	case 3:
		x = 103
		if val > 100 {
			fmt.Println("val")
		}
	default:
		x = 104
		fmt.Println("default")
	}
	return x
}
