package xgo_cache

import (
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

const (
	first  = iota
	second = iota
	th4    = iota
	th5    = iota
)

type (
	s struct {
		i   int
		arr [10]int
	}
)

func cacheWork() {
	cache := gocache.New(20*time.Second, 1*time.Second)
	sval := s{}
	cache.Add("1", &sval, 5*time.Second)
	sval.arr[1] = 100
	rval, _ := cache.Get("1")
	fmt.Println(rval)
}
