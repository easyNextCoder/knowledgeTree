package skiplist2

import (
	"fmt"
	"math/rand"
	"time"
)

var maxLevel int = 32
var pFactor float32 = 0.25

func getRandom() bool {

	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(50) > 25
	return res
}

func getRandomInt() int {
	rand.Seed(time.Now().UnixNano())

	initHeight := 1
	for initHeight <= maxLevel && rand.Float32() < pFactor {
		initHeight++
	}
	return initHeight
}

var maxLen int = 32

type Node struct {
	forward []*Node
	val     int
}

type Skiplist struct {
	head  *Node
	level int
}

func Constructor() Skiplist {
	//res := getRandomInt()
	return Skiplist{head: &Node{
		forward: make([]*Node, maxLen),
		val:     -1,
	},
		level: 0}
}

func (this *Skiplist) Search(target int) bool {
	curr := this.head
	for i := this.level; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	return curr != nil && curr.forward[0] != nil && curr.forward[0].val == target
}
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func (this *Skiplist) Add(num int) {
	update := make([]*Node, maxLen)
	curr := this.head
	nowLevel := getRandomInt()
	for i := max(this.level-1, nowLevel-1); i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}

	newNode := &Node{
		forward: make([]*Node, nowLevel),
		val:     num,
	}

	for i := nowLevel - 1; i >= 0; i-- {
		fmt.Println(update[i].forward[i])
		newNode.forward[i] = update[i].forward[i]

		update[i].forward[i] = newNode
		fmt.Println(update[i].forward[i])
	}
	this.level = max(this.level, nowLevel)

}

func objPrint(sk Skiplist) {

	for i := sk.level - 1; i >= 0; i-- {
		curr := sk.head
		fmt.Print(curr.val, " ")
		for curr.forward[i] != nil {
			fmt.Print(curr.forward[i].val, "  ")
			curr = curr.forward[i]
		}
		fmt.Println("\n=========================")
	}
}

func (this *Skiplist) Erase(num int) bool {
	update := make([]*Node, this.level)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	if curr.forward[0] == nil || curr.forward[0].val != num {
		return false
	}

	for i, _ := range update {
		if update[i].forward[i] != nil && update[i].forward[i].val == num {
			update[i].forward[i] = update[i].forward[i].forward[i]
		}
	}
	return true
}

func runSkipList2() {
	obj := Constructor()

	//param_1 := obj.Search(target)
	obj.Add(1)
	obj.Add(10)
	obj.Add(5)
	obj.Add(6)

	obj.Add(100)
	obj.Add(100)
	obj.Add(40)
	obj.Add(35)
	obj.Add(102)
	obj.Add(103)
	obj.Add(104)
	obj.Add(106)
	obj.Add(1090)
	obj.Add(106)
	obj.Add(107)
	obj.Add(109)
	obj.Add(1034)
	obj.Add(112)
	obj.Add(123)
	obj.Add(145)
	obj.Add(167)
	obj.Add(189)
	obj.Add(199)
	obj.Add(156)
	objPrint(obj)
	//
	fmt.Println(obj.Search(1))
	fmt.Println(obj.Search(200))
	fmt.Println(obj.Search(109))
	fmt.Println(obj.Erase(109))
	fmt.Println(obj.Search(109))
	fmt.Println(obj.Erase(100))
	fmt.Println(obj.Search(100))
	fmt.Println(obj.Erase(999999))
	//fmt.Println(obj.Search(107))
	//fmt.Println(obj.Search(145))
	//fmt.Println(obj.Search(123))
	//fmt.Println(obj.Erase(100))
	//
	//fmt.Println(obj.Search(1999))
	//fmt.Println(obj.Search(5))
	//fmt.Println(obj.Search(123))
}
