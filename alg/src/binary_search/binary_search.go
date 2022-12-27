package binary_search

import "fmt"

func myLowerBound(data []int, k int) int {
	start := 0
	last := len(data)
	for start < last {
		mid := (start + last) / 2
		if data[mid] >= k {
			last = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func myUpperBound(data []int, k int) int {
	start := 0
	last := len(data)
	for start < last {
		mid := (start + last) / 2
		if data[mid] > k {
			last = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func once() {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 7, 8, 10}
	//lowBounde->index=4
	//upperBound->index=8
	//所以最终找的是一个左闭右开的区间
	fmt.Println(myLowerBound(arr, 5))
	fmt.Println(myUpperBound(arr, 5))

}
