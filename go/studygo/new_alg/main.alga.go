package main

import (
	"net/http"
	_ "net/http/pprof"
	j "studygo/new_alg/alga"
)

func main() {
	var a j.AlgX

	go func() {
		http.ListenAndServe("localhost:6061", nil)
	}()

	a.Work()
}
