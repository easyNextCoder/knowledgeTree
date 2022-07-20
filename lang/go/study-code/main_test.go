package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}
	var oriArr = [4]int{1, 2, 3, 4}

	fmt.Println(oriArr)
	//testExchangeArr(oriArr)
	fmt.Println(oriArr)

	arr := [10]int{1, 2, 3}
	arr1 := arr[:5]
	arr2 := arr[5:8]
	arr3 := arr[4:]
	fmt.Println("arr")
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println("arr1")
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(arr)
	fmt.Println(arr1)
	arr1x := append(arr1, 99)
	fmt.Println(cap(arr1x))
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	arr1x = append(arr1x, 99)
	fmt.Println(cap(arr1x))
	arr1[1] = 100
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr1x)
	fmt.Println("arr2")
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))
	fmt.Println("arr3")
	fmt.Println(len(arr3))
	fmt.Println(cap(arr3))

}
func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
