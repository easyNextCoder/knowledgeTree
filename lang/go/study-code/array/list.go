package array

import "fmt"

var s [6]string

func testArray() {
	fmt.Println("testArray:")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"
	s[5] = "f"

	s1 := [6]string{"1", "2", "3", "4", "5", "6"}

	fmt.Println(s[1:4])
	var s1Copy = s1[1:4]
	s2Copy := s1[2:5]
	fmt.Println(s1[1:4], s1Copy, s2Copy)

	//对应的内建函数
	//make
	//len
	//caps
	//append
}
