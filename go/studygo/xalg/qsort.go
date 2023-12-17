package xalg

func quickSort20231206(arr []int, start, end int) {
	if end-start < 2 {
		return
	}

	mid := (start + end) / 2
	swap(start, mid, arr)
	idx := start + 1
	for i := start + 1; i < end; i++ {
		if arr[i] < arr[start] {
			swap(i, idx, arr)
			idx++
		}
	}
	swap(start, idx-1, arr)
	quickSort20231206(arr, start, idx-1)
	quickSort20231206(arr, idx, end)

}
