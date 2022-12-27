package xsort

import "fmt"

func swap2(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

func quickSort2(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	swap2(arr, start, mid)
	l, r := start, end

	for l < r {
		for l < r && arr[r] > arr[start] {
			r--
		}
		for l < r && arr[l] < arr[start] {
			l++
		}
		if l < r {
			swap2(arr, l, r)
		}
	}

	swap2(arr, start, l)
	quickSort2(arr, start, l)
	quickSort2(arr, l+1, end)
}

func runQuickSort2() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	quickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
