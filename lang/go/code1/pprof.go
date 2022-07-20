package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func PprofServer() {
	log.Println(http.ListenAndServe(":6060", nil))
}

func workOnMain() {
	i := 0
	for {

		fmt.Println("main working once ", i)
		time.Sleep(1000 * time.Millisecond)
		i++
	}
}

//go func() {
//	for {
//		fmt.Println("work once")
//		time.Sleep(500 * time.Millisecond)
//	}
//}()
//go PprofServer()
//workOnMain()
