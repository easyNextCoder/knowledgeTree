package xalg

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func qsort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	tmp := arr[mid]
	arr[mid] = arr[l]
	arr[l] = tmp

	left := l
	right := r
	for left < right {
		for left < right && arr[right] >= arr[l] {
			right--
		}
		for left < right && arr[left] <= arr[l] {
			left++
		}

		if left != right {
			tmp1 := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp1
		}
	}
	tmp2 := arr[left]
	arr[left] = arr[l]
	arr[l] = tmp2
	fmt.Println(left)
	qsort(arr, l, left-1)
	qsort(arr, left+1, r)
}

func testSort() {

	arr := []int{6, 5, 7, 8, 9, 10, 12, 13}
	qsort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
