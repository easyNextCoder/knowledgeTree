package xmap

import (
	"fmt"
	"time"
)

func mapInitBench() {

	s := time.Now()

	for i := 0; i < 100000000; i++ {
		v := make(map[int]int)
		v[1] = 1
	}

	e := time.Now()
	for i := 0; i < 100000000; i++ {
		var x [100]bool
		x[1] = true
	}

	f := time.Now()

	fmt.Println(e.Sub(s).Milliseconds(), f.Sub(e).Milliseconds())
}
