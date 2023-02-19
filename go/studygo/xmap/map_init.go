package xmap

import "fmt"

func do_something() {
	fmt.Println("do something")
}

func mapInit() {
	//wrong code
	boolMap := make(map[int]bool)
	boolMap[1] = true
	boolMap[2] = true
	third := boolMap[3] // third will be false
	fmt.Println("访问不存在的key", third)

	//useage
	if third2, ok := boolMap[3]; ok {
		fmt.Println(third2)
	} else {
		//third has value
	}

	// other useage
	if boolMap[3] { // false or can't find key
		do_something()
	}

	xr, xv := boolMap[3]
	xt := boolMap[3]
	fmt.Println("访问不存在的key，返回的结果以及格式", xr, xv, xt)

	intMap := map[int]int{}
	yr, yv := intMap[3]
	yt := intMap[3]
	fmt.Println("int类型map，访问不存在的key，返回的结果以及格式", yr, yv, yt)

	structMap := map[int]struct{}{}
	zr, zv := structMap[3]
	zt := structMap[3]
	fmt.Println("struct类型map，访问不存在的key，返回的结果以及格式", zr, zv, zt)

	structPointerMap := map[int]*struct{}{}
	ar, av := structPointerMap[3]
	at := structPointerMap[3]
	fmt.Println("指针类型map，访问不存在的key，返回的结果以及格式", ar, av, at, at == nil)

}
