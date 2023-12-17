package xalg

func AscChecker(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if !(arr[i] >= arr[i-1]) {
			return false
		}
	}

	return true
}

func DescChecker(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}

func MinHeapCheck(arr []int, root int) bool {
	l, r := root*2+1, root*2+2

	lok, rok := true, true

	if l < len(arr) {
		if arr[root] <= arr[l] {
			lok = MinHeapCheck(arr, l)
		} else {
			return false
		}
	}
	if r < len(arr) {
		if arr[root] <= arr[r] {
			rok = MinHeapCheck(arr, r)
		} else {
			return false
		}
	}

	return lok && rok
}
