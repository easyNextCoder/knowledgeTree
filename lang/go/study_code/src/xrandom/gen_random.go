package xrandom

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var randItem = struct {
	sync.Mutex
	*rand.Rand
}{
	Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

func RandRange(i, j int) int {
	min := i
	max := j
	if min > max {
		min, max = max, min
	}
	if (max - min) <= 0 {
		panic("invalid argument to randrange max cant equal min")
	}
	randItem.Lock()
	defer randItem.Unlock()
	x := randItem.Intn(max-min) + min
	return x
}

func genRand() {
	for i := 0; i < 100; i++ {
		fmt.Println(RandRange(0, 2))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(RandRange(0, 100))
	}

}
