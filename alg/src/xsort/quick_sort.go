package xsort

import "fmt"

//215. 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/
func quickSort(nums []int, start, end int, k int) {
	if end-start+1 < 2 {
		return
	}
	mid := (start + end) / 2
	swap(nums, start, mid)
	l, r := start, end
	for l < r {
		for l < r && nums[r] >= nums[start] {
			r--
		}
		for l < r && nums[l] <= nums[start] {
			l++
		}

		if l < r {
			swap(nums, l, r)
		}
	}

	swap(nums, start, l)
	//if k == l {
	//	return
	//} else if k < l+1 {
	//	quickSort(nums, start, l-1, k)
	//} else {
	//	quickSort(nums, l+1, end, k)
	//}

	quickSort(nums, start, l-1, k)
	quickSort(nums, l+1, end, k)
}

func findKthLargest2(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, len(nums)-k)
	fmt.Println(nums)
	return nums[len(nums)-k]
}
