package xalg

import (
	"fmt"
	"sort"
)

func sortSliceWork() {
	type pair struct {
		card   int
		profit int
		score  int
	}
	cp := make([]pair, 0)
	for _, card := range []int{1, 2, 3} {

		cp = append(cp, pair{card: card, profit: card + 100, score: card})
	}
	cp[0].profit = 102
	sort.Slice(cp, func(i, j int) bool {
		if cp[i].profit == cp[j].profit {
			return cp[i].score > cp[j].score
		} else {
			return cp[i].profit > cp[j].profit
		}
	})
	fmt.Println(cp)
}
