package xalg

import "math/rand"

func swap(a, b int, arr []int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

func GenRandomArr(num, start, end int) []int {
	ret := []int{}

	for i := 0; i < num; i++ {
		ret = append(ret, rand.Int()%(end-start))
	}

	return ret
}
