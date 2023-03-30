package rules

import (
	"fmt"
	"math/big"
)

// 找到最小的比N大的，二进制有相同个数的1的数
func NextSame1(n int) int {
	var x = n & -n
	var t = n + x
	if x == 0 {
		panic(fmt.Errorf("x is zero %d", n))
	}
	return t | ((n^t)/x)>>2
}

// N个取M个的次数
func Combine(n, m int) int {
	var up = big.NewInt(1).MulRange(int64(n-m)+1, int64(n))
	var down = big.NewInt(1).MulRange(1, int64(m))
	return int(up.Div(up, down).Int64())
}

func fillCard(allChoice []Card, hashCode int, fillList []Card) {
	var putPos = 0
	var idx = 0
	for j := hashCode; j > 0; j = j >> 1 {
		if j&1 > 0 {
			fillList[putPos] = allChoice[idx]
			putPos += 1
		}
		idx += 1
	}
}
func fixMinMax(cardCount int, min, max int) (int, int) {
	for _, v := range [2]int{min, max} {
		if v < 0 || v > cardCount {
			panic(fmt.Errorf("bad min max %d %d", min, max))
		}
	}
	if min < max {
		return min, max
	} else {
		return max, min
	}
}

func SubSet(cards []Card, count int) [][]Card {
	return SubSetRange(cards, count, count)
}

func SubSetRange(cards []Card, min, max int) [][]Card {
	min, max = fixMinMax(len(cards), min, max)
	var subCount = 0
	var cardCount = 0
	for i := min; i <= max; i++ {
		var count = Combine(len(cards), i)
		subCount += count
		cardCount += count * i
	}
	var result = make([][]Card, subCount)
	var mempool = make([]Card, cardCount)
	var resultPos = 0
	for k := min; k <= max; k++ {
		var start = (1 << k) - 1
		var end = 1<<len(cards) - 1<<(len(cards)-k)
		//注意bit的最大值为(1<<n) - (1<<(n-k)即选择n的前k个
		for bit := start; bit <= end; bit = NextSame1(bit) {
			fillCard(cards, bit, mempool[0:k])
			result[resultPos] = mempool[0:k]
			resultPos += 1
			mempool = mempool[k:]
		}
	}
	return result
}
