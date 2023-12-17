package xalg

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {

	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}
	n = 50
	trips = [][]int{{1, 3}}
	price = make([]int, 50, 50)
	price[0], price[1], price[2], price[3], price[4] = 100, 1, 100, 1, 1

	var count = make([]int, 51, 51)

	var m = make([][]bool, n, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, 51, 51)
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

	var workDp func(parent, node int) []int

	workDp = func(parent, node int) []int {
		//fmt.Println("workDp", parent, node)
		res := []int{count[node] * price[node], count[node] * price[node] / 2}

		for k, v := range m[node] {

			if !v {
				continue
			}

			if k == parent {
				continue
			}

			rest := workDp(node, k)

			res[0], res[1] = res[0]+min(rest[0], rest[1]), res[1]+rest[0]

		}
		return res
	}

	res := workDp(-1, 0)

	//fmt.Println(count, res)
	return min(res[0], res[1])
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

func minimumTotlPrice(n int, edges [][]int, price []int, trips [][]int) int {

	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}
	n = 50
	trips = [][]int{{1, 3}}
	price = make([]int, 50, 50)
	price[0], price[1], price[2], price[3], price[4] = 100, 1, 100, 1, 1

	next := make([][]int, n)
	for _, edge := range edges {
		next[edge[0]] = append(next[edge[0]], edge[1])
		next[edge[1]] = append(next[edge[1]], edge[0])
	}

	count := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(node, parent, end int) bool {
		if node == end {
			count[node]++
			return true
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			if dfs(child, node, end) {
				count[node]++
				return true
			}
		}
		return false
	}
	for _, trip := range trips {
		dfs(trip[0], -1, trip[1])
	}

	var dp func(int, int) []int
	dp = func(node, parent int) []int {
		fmt.Println("dp", parent, node)
		res := []int{
			price[node] * count[node], price[node] * count[node] / 2,
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			v := dp(child, node)
			// node 没有减半，因此可以取子树的两种情况的最小值
			// node 减半，只能取子树没有减半的情况
			res[0], res[1] = res[0]+min(v[0], v[1]), res[1]+v[0]
		}
		return res
	}
	res := dp(0, -1)

	fmt.Println(count, res)
	return min(res[0], res[1])

}
