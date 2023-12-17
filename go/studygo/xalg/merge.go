package xalg

func mergeSort(a, b int, arr []int) { //panic
	if !(a < b) {
		return
	}

	mid := (a + b) / 2
	mergeSort(a, mid, arr)
	mergeSort(mid+1, b, arr)

	tmp := make([]int, b-a+1)
	idxt, idxa, idxb := 0, a, mid+1

	for idxa <= mid || idxb <= b {

		minVal := 0
		if idxa > mid {
			minVal = arr[idxb]
			idxb++
		} else if idxb > b {
			minVal = arr[idxa]
			idxa++
		} else {

			if arr[idxa] < arr[idxb] {
				minVal = arr[idxa]
				idxa++
			} else {
				minVal = arr[idxb]
				idxb++
			}
		}

		tmp[idxt] = minVal
		idxt++

	}
	for i := a; i <= b; i++ {
		arr[i] = tmp[i-a]
	}
}
