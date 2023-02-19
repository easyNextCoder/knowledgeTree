package xconst

import "fmt"

const (
	X = iota
	Y = iota
	YY
	Z = iota
	ZA
	ZB = 100
	ZC = iota
	ZZ
	ZZZ = iota
)

func IotaWork() {
	fmt.Println(X, Y, YY, Z, ZA, ZB, ZC, ZZ, ZZZ)
	//0 1 2 3 4 100 6 7 8
}
