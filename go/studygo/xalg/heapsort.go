package xalg

import "fmt"

func justifyHeap20231130(arr []int, father, end int) {
	lson := father*2 + 1
	rson := father*2 + 2
	if !(lson < end) {
		return
	}

	if arr[father] > arr[lson] {
		swap(father, lson, arr)
		justifyHeap20231130(arr, lson, end)
	}

	if rson < end && arr[father] > arr[rson] {
		swap(father, rson, arr)
		justifyHeap20231130(arr, rson, end)
	}
}

func heapSort20231130(arr []int) {

	for i := len(arr)/2 - 1; i >= 0; i-- {
		justifyHeap20231130(arr, i, len(arr))
	}

	checkOk := MinHeapCheck(arr, 0)

	fmt.Printf("小根堆建立完成 checkOk(%t) arr(%+v)\n", checkOk, arr)

	res := []int{}
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[0])
		arr[0] = arr[len(arr)-1-i]
		justifyHeap20231130(arr, 0, len(arr)-1-i)
	}

	fmt.Println("排序结果:", res, AscChecker(res))

	arr = res
}
