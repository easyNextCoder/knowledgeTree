package arr

import "fmt"

func sub(a [10]int) {
	a[0] = 100
	fmt.Printf("%p %v\n", &a, a)
}

func arrCopyWork() {
	a := [10]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("before %p %v\n", &a, a)
	sub(a)
	fmt.Printf("%p %v\n", &a, a)
}
