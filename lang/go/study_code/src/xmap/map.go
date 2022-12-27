package xmap

import (
	"fmt"
	"include/data"
	"strings"
	"time"
)

func WordCount(s string) map[string]int {
	con := make(map[string]int)
	vs := strings.Fields(s)
	for _, ss := range vs {
		v, ok := con[ss]
		if ok {
			con[ss] = v + 1
		} else {
			con[ss] = 1
		}
	}
	return con
}

var xmp map[string]string

func testMap() {
	xmp["name"] = "your"
	fmt.Println(xmp["name"])
	var res = WordCount("fdf i am the king, who are you, and are you?")
	fmt.Println(res)
}

func DeleteInTraverse() {
	mp := map[string]int{
		"hello": 0,
		"world": 1,
		"walk":  2,
		"away":  3,
	}
	arr := []string{"world", "walk", "hello", "away"}
	for k, _ := range mp {
		for _, aval := range arr {
			if k == aval {
				delete(mp, k)
				fmt.Println("delete in mp", k)
			}
		}
	}
	fmt.Println(mp)
}

func UninitializedMapRead() {
	var mp map[string]int
	fmt.Println(mp["hello"])
	fmt.Println(mp["word"])
	//读取空的map没有问题，也不会进行插入操作
}

func UninitializedMapWrite() {
	var mp map[string]int
	fmt.Println(mp["word"])
	//读取不会报错
	mp["word"] = 1
	//直接写入空的map会有问题
	//err:assignment to entry in nil map
}

func GoroutineReadUninitializedMap(mp map[string]int, key string) {
	for i := 0; i < 100; i++ {
		fmt.Println(key, i, mp[key])
	}
}

func ManyGoroutineReadUninitializedMap() {
	var mp map[string]int
	go GoroutineReadUninitializedMap(mp, "hello")
	go GoroutineReadUninitializedMap(mp, "word")
	defer GoroutineReadUninitializedMap(mp, "xyk")
	time.Sleep(200 * time.Millisecond)
}

func GoroutineWriteUninitializedMap(mp map[string]int, key string) {
	for i := 0; i < 100; i++ {
		mp[key] = i
		fmt.Println(key, i)
	}
}

func ManyGoroutineWriteUninitializedMap() {
	var mp map[string]int
	go GoroutineWriteUninitializedMap(mp, "hello")
	go GoroutineWriteUninitializedMap(mp, "word")
	time.Sleep(200 * time.Millisecond)
}

func GoroutineReadInitializedMap(mp map[string]int, key string) {
	for i := 0; i < 100; i++ {
		k, v := mp[key]
		if v {
			k++
		}
	}
}

func ManyGoroutineReadInitializedMap() {
	mp := make(map[string]int)
	go GoroutineReadInitializedMap(mp, "goa")
	go GoroutineReadInitializedMap(mp, "gob")
	go GoroutineReadInitializedMap(mp, "goc")
	go GoroutineReadInitializedMap(mp, "god")
	go GoroutineReadInitializedMap(mp, "goe")
	go GoroutineReadInitializedMap(mp, "gof")
	go GoroutineReadInitializedMap(mp, "gog")

	time.Sleep(200 * time.Millisecond)
}

func GoroutineWriteInitializedMap(mp map[string]int, key string) {
	for i := 0; i < 100; i++ {
		mp[key] = i
	}
}

func ManyGoroutineWriteInitializedMap() {
	mp := make(map[string]int)
	go GoroutineWriteInitializedMap(mp, "goa")
	go GoroutineWriteInitializedMap(mp, "gob")
	go GoroutineWriteInitializedMap(mp, "goc")
	go GoroutineWriteInitializedMap(mp, "god")
	time.Sleep(200 * time.Millisecond)
	//fatal error: concurrent map writes
}

func ManyGoroutinesWriteReadInitializedMap() {
	mp := make(map[string]int)
	go GoroutineWriteInitializedMap(mp, "goa")
	go GoroutineReadInitializedMap(mp, "goa")
	go GoroutineWriteInitializedMap(mp, "gob")
	go GoroutineReadInitializedMap(mp, "gob")
	go GoroutineWriteInitializedMap(mp, "goc")
	go GoroutineReadInitializedMap(mp, "goc")
	time.Sleep(200 * time.Millisecond)
	//fatal error: concurrent map read and map write
	//fatal error: concurrent map writes
}

type node struct {
	val1 int
	arr  []int
}

func UseValueAfterDelete() {
	//测试删除之前获取value，删除这个key之后这个value是否还可用
	//结果：实际上是可用的
	mp := make(map[int]*node)
	arrp := make([]int, 0)
	arrp = append(arrp, 1)
	arrp = append(arrp, 2)
	arrp = append(arrp, 3)

	mp[1] = &node{val1: 1, arr: arrp}

	arrb := mp[1]
	fmt.Println("originMp", mp)
	delete(mp, 1)
	fmt.Println("deletedMp", mp)
	fmt.Println(arrb)

}

func SimpleInitMap() {
	mp := make(map[string]bool)
	mp["workname"] = false
	mp["workhello"] = true
	mp["workfind"] = true
	delete(mp, "hello")
	fmt.Println(mp)

	mp2 := make(map[int]bool)
	fmt.Println(mp2[1])

	fmt.Println("see what have run", data.X1)
}

func mapParaToFunc() {
	a := make(map[int]int)
	a[1] = 0
	a[2] = 0
	runWithFunc(a)
	fmt.Println(a)
}

func runWithFunc(a map[int]int) {
	for i, _ := range a {
		a[i] = 100
	}
}
