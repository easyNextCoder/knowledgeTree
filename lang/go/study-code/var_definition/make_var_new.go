package var_definition

import "fmt"

func vart() {
	var p []int

	b := []int{}

	fmt.Println(p, b)
}

func maket() {
	p := new([]int)

	b := []int{}

	fmt.Println(p, b)
}

func newt() {
	p := make([]int, 0)

	b := []int{}

	fmt.Println(p, b)
}

var tag [51][51]int
var ISOLATED int = -1
var ISOLATED1 int = -2
var INFECTED int = 1
var GOOD int = 0
var ACCESSED int = 2
var ACCESSED1 int = 3
var finalCount int = 0

func check(isInfected [][]int) bool {
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == GOOD || isInfected[i][j] == ISOLATED1 {
				return true
			}
		}
	}
	return false
}

func infect(isInfected [][]int, x, y int) {
	if !(x < len(isInfected) && x >= 0 && y < len(isInfected[0]) && y >= 0) {
		return
	}
	f := func(isInfected [][]int, x, y int) bool {
		if x < len(isInfected) && y < len(isInfected[0]) && isInfected[x][y] == GOOD || isInfected[x][y] == ISOLATED1 {
			isInfected[x][y] = ACCESSED
			return true
		}
		return false
	}
	if f(isInfected, x, y) {
		return
	} else if isInfected[x][y] == INFECTED || isInfected[x][y] == ISOLATED1 {
		isInfected[x][y] = ACCESSED1
		infect(isInfected, x+1, y)
		infect(isInfected, x-1, y)
		infect(isInfected, x, y+1)
		infect(isInfected, x, y-1)
	}

}

func preGrow(isInfected [][]int) (i, j int) {
	isInfectedCopy := make([][]int, len(isInfected))
	for i, _ := range isInfectedCopy {
		isInfectedCopy[i] = make([]int, len(isInfected[0]))
	}
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			isInfectedCopy[i][j] = isInfected[i][j]
		}
	}

	retx, rety := -1, -1
	max := 0
	clr := func(isInfected [][]int) {
		for i := 0; i < len(isInfected); i++ {
			for j := 0; j < len(isInfected[i]); j++ {
				isInfected[i][j] = isInfectedCopy[i][j]
			}
		}
	}
	cnt := func(isInfected [][]int, i, j int) int {
		count := 0
		for i := 0; i < len(isInfected); i++ {
			for j := 0; j < len(isInfected[i]); j++ {
				if isInfected[i][j] == ACCESSED {
					count++
					isInfected[i][j] = ACCESSED1
				}
			}
		}
		return count
	}
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == INFECTED || isInfected[i][j] == ISOLATED1 {

				infect(isInfected, i, j)
				tmpCnt := cnt(isInfected, i, j)
				if max < tmpCnt {
					max = tmpCnt
					retx, rety = i, j
				}
			}
		}
	}
	clr(isInfected)
	return retx, rety
}

func grow(isInfected [][]int) bool {
	fmt.Println("v", isInfected)
	count := 0
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == INFECTED {
				count++
			}
		}
	}
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == INFECTED {
				infect(isInfected, i, j)
			}
		}
	}
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == ACCESSED || isInfected[i][j] == ACCESSED1 {
				isInfected[i][j] = INFECTED
				count--
			}
		}
	}
	fmt.Println("v", isInfected)
	if count == 0 {
		return false
	}
	return true
}

func isolateOne(isInfected [][]int, x, y int) {

	if !(x < len(isInfected) && x >= 0 && y < len(isInfected[0]) && y >= 0) {
		return
	}
	f := func(isInfected [][]int, x, y int) bool {
		if x < len(isInfected) && y < len(isInfected[0]) && (isInfected[x][y] == GOOD || isInfected[x][y] == ISOLATED1) {
			finalCount++
			isInfected[x][y] = ISOLATED1
			return true
		}
		return false
	}
	if f(isInfected, x, y) {
		return
	} else if isInfected[x][y] == INFECTED {
		isInfected[x][y] = ISOLATED
		isolateOne(isInfected, x+1, y)
		isolateOne(isInfected, x-1, y)
		isolateOne(isInfected, x, y+1)
		isolateOne(isInfected, x, y-1)
	}
}

func isolate(isInfected [][]int) bool {
	res := false
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == INFECTED {
				isolateOne(isInfected, i, j)
				res = true
				return res
			}
		}
	}
	return res
}

func containVirus(isInfected [][]int) int {
	for check(isInfected) {
		x, y := preGrow(isInfected)
		fmt.Println(x, y)
		isolateOne(isInfected, x, y)

		res := grow(isInfected)
		fmt.Println(isInfected)
		if !res {
			fmt.Println(res)
			break
		}
	}
	fmt.Println(isInfected)
	count := 0
	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == ISOLATED1 {
				count++
			}
		}
	}
	return finalCount
}

//[[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]
// [[1,1,1,0,0,0,0,0,0],
//  [1,0,1,0,1,1,1,1,1],
//  [1,1,1,0,0,0,0,0,0]]

// [[1 1 1 0 -2 -2 -2 -2 -2] [1 0 1 -2 -1 -1 -1 -1 -1] [1 1 1 0 -2 -2 -2 -2 -2]]
// [[1 1 1 1 -2 -2 -2 -2 -2] [1 1 1 -2 -1 -1 -1 -1 -1] [1 1 1 1 -2 -2 -2 -2 -2]]
// [[1 1 1 1 -2 -2 -2 -2 -2] [1 1 1 -2 -1 -1 -1 -1 -1] [1 1 1 1 -2 -2 -2 -2 -2]]
// [[1 1 1 1 -2 -2 -2 -2 -2]
//  [1 1 1 -2 -1 -1 -1 -1 -1]
//  [1 1 1 1 -2 -2 -2 -2 -2]]


