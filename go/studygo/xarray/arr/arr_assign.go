package arr

import "fmt"

func arrAssignWork() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Printf("a %p %v\n", &a, a)
	fmt.Printf("b %p %v\n", &b, b)
}
