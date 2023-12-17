package files

import (
	"fmt"
	"time"
)

type Son struct {
	p []int
}
type TestStruct struct {
	son  *Son
	id   int
	name string
}

func mainGo() {

	var t = TestStruct{
		son:  &Son{p: []int{1, 2, 3}},
		id:   0,
		name: "xiaohong",
	}

	fmt.Println("mainGo", t)

	sonGo(&t)

}

func sonGo(x *TestStruct) {

	time.Sleep(time.Second * 2)
	fmt.Println("sonGo", x.son.p)
	grandSonGo(x)
}

func grandSonGo(x *TestStruct) {

	time.Sleep(time.Second * 2)
	fmt.Println("grandSonGo", x.son.p)
}
