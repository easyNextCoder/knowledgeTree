package xsort

//215. 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/

func maxHeapify(nums []int, a int, heapSize int) {
	l, r, largest := a*2+1, a*2+2, a
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != a {
		swap(nums, a, largest)
		maxHeapify(nums, largest, heapSize)
	}
}

func findKthLargest(nums []int, k int) int {

	heapSize := len(nums)
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
	res := nums[0]
	for i := 0; i < k; i++ {
		res = nums[0]
		nums[0] = nums[heapSize-1]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}

	return res
}
