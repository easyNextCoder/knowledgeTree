package xsort

func swap(nums []int, ia, ib int) {
	tmp := nums[ia]
	nums[ia] = nums[ib]
	nums[ib] = tmp
}
