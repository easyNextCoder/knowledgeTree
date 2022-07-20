package main

import "fmt"

func appendArray() {
	a, b := []int{1, 2, 3}, []int{4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
}

func getKindsOfParas(v ...interface{}) {
	for _, v := range v {
		switch vv := v.(type) {
		case int:
			fmt.Println("int", vv)
		case string:
			fmt.Println("string", vv)
		case map[string]int:
			fmt.Println("map[string]int", vv)
		case []int:
			fmt.Println("[]int", vv)
			for index, val := range vv {
				fmt.Println(index, val)
			}
		}
	}
}

func workOnGetKindsOfParas() {
	a := 1
	b := "name"
	c := map[string]int{}
	d := []int{1, 2, 3, 4}
	getKindsOfParas(a, b, c, d)
}
