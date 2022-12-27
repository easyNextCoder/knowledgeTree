package dynamicProgramming

import "fmt"

var minVal int = 1000000
var tmpMinVal int = 1000000
var cache map[int]int = map[int]int{}
var cache2 map[int]int = map[int]int{}

func work(nums1 []int, nums2 []int, i int, cnt int) (int, int) { //不交换， 交换

	if i == 0 {
		return 0, 1
	}

	if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
		a2, b2 := work(nums1, nums2, i-1, cnt+1)

		cache[i] = a2
		cache2[i] = b2 + 1
	}

	if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {

		a1, b1 := work(nums1, nums2, i-1, cnt+1)

		cache[i], cache2[i] = min(cache[i], b1), min(cache2[i], a1+1)

	}
	return cache[i], cache2[i]
}

//11.52

func swap(nums []int, nums2 []int, index int) {

	tmp := nums[index]
	nums[index] = nums2[index]
	nums2[index] = tmp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSwap(nums1 []int, nums2 []int) int {

	// cache[len(nums1)-1] = 0
	a, b := work(nums1, nums2, len(nums1)-1, 0)
	fmt.Println(a, b, cache, nums1, nums2)
	return cache[0]
}

func run() {
	a := []int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}
	b := []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}

	for i, _ := range a {
		cache[i] = 10000000
		cache2[i] = 10000000
	}
	fmt.Println("main", cache[9], cache2[9], cache, cache2)

	//a := []int{0, 7, 8, 7, 10, 11, 12, 13, 19, 20}
	//b := []int{4, 4, 5, 10, 11, 14, 15, 16, 17, 18}

	//a := []int{0, 3, 5, 8, 9}
	//b := []int{2, 1, 4, 6, 9}
	//a := []int{1, 3, 5, 4}
	//b := []int{1, 2, 3, 7}

	minSwap(a, b)
}

// [0,4,8,10,10,11,12,13,19,18]
// [4,7,5,7,11,14,15,16,17,20]
