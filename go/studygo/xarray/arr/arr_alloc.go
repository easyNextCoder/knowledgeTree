package arr

import "fmt"

type X struct {
	int
	arr  [10]int
	arrp []int
}

var xarr [4]*X

func load() {
	for i := range xarr {
		xarr[i] = new(X)
	}
}

func arrAllocWork() {
	load()
	for i, v := range xarr {
		v.arrp = append(v.arrp, i)
		fmt.Printf("this is %d %+v\n", i, v)
	}
}

var n = 50
var count = make([]int, 50, 50)
var prices = make([]int, 50, 50)
var m = make([][]bool, n, n)
var edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}
var trips = [][]int{{1, 3}}

func work() {

	for i := 0; i < n; i++ {
		m[i] = make([]bool, 50, 50)
	}

	for _, v := range edges {
		a, b := v[0], v[1]
		m[a][b] = true
		m[b][a] = true
	}

	for _, v := range trips {
		start, end := v[0], v[1]
		workDfs(m, count, -1, start, end)
	}

	res := workDp(-1, 0)
	return min(res[0], res[1])
	fmt.Println("This is final count:", count)

}

func workDp(parent, node int) []int {

	res := []int{count[node] * prices[node], count[node] * prices[node] / 2}

	for k, v := range m[node] {

		if !v {
			continue
		}

		rest := workDp(node, k)

		res[0], res[1] = rest[0]+min(rest[0], rest[1]), rest[1]+res[1]

	}
	return res
}

func workDfs(m [][]bool, count []int, parent, node, end int) bool {
	if node == end {
		count[node]++
		return true
	}

	for i, v := range m[node] {
		if !v {
			continue
		}

		if i == parent {
			continue
		}

		if workDfs(m, count, node, i, end) {
			count[node]++
			return true //如果少这一行，只有终点和last终点被染色
		}

	}

	return false
}
