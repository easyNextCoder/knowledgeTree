package xalg

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	val int

	level int

	up *Node

	down *Node

	left *Node

	right *Node
}

var maxLevel int = 32

func getRandom() bool {

	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(50) > 25
	fmt.Println("getRandom:", res)
	return res
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {

	return Skiplist{head: &Node{
		val:   -1,
		level: 1,
		up:    nil,
		down:  nil,
		left:  nil,
		right: nil,
	}}

}

func (this *Skiplist) Search(target int) bool {

	tmp := this.head

	for tmp != nil {
		if tmp.val == target {
			return true
		} else {
			if tmp.right != nil {
				if target == tmp.right.val {
					return true
				} else if target < tmp.right.val {
					tmp = tmp.down
				} else {
					tmp = tmp.right
				}
			} else {
				tmp = tmp.down
			}
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {

	tmp := this.head
	var pre *Node
	myStack := make([]*Node, 0)
	for tmp != nil {
		pre = tmp
		myStack = append(myStack, pre)
		if tmp.right == nil {
			tmp = tmp.down
		} else {
			if num <= tmp.right.val {
				tmp = tmp.down
			} else {
				tmp = tmp.right
			}
		}
	}

	newNode := &Node{
		val:   num,
		level: 1,
		up:    nil,
		down:  nil,
		left:  pre,
		right: pre.right,
	}
	if pre.right != nil {
		pre.right.left = newNode
	}

	pre.right = newNode
	fmt.Println("zzzz", pre, num)
	fmt.Println(pre, "add node", newNode)
	if getRandom() {
		for i := 2; i < maxLevel && getRandom(); i++ {
			for myStack[len(myStack)-1].level < i {
				myStack = myStack[:len(myStack)-1]
				if len(myStack) == 0 {
					break
				}
			}
			fmt.Println(myStack)
			if len(myStack) == 0 {
				//这里要更新head了
				newHead := &Node{
					val:   this.head.val,
					level: i,
					up:    nil,
					down:  this.head,
					left:  nil,
					right: nil,
				}
				this.head.up = newHead
				this.head = newHead
				myStack = append(myStack, this.head)

			}

			beforeNewNode := myStack[len(myStack)-1]
			newNode.up = &Node{
				val:   num,
				level: i,
				up:    nil,
				down:  newNode,
				left:  beforeNewNode,
				right: beforeNewNode.right,
			}
			newNode = newNode.up
			if beforeNewNode.right != nil {
				beforeNewNode.right.left = newNode
			}
			beforeNewNode.right = newNode
		}

	}

}

func (this *Skiplist) Erase(num int) bool {
	target := num
	tmp := this.head

	del := func(tmp *Node) {
		fmt.Println("yyyyyyyyy", tmp, tmp.left)
		for tmp != nil {
			tmpBack := tmp
			pre := tmp.left
			fmt.Println("XXxxxxxx", pre, tmp)
			next := tmp.right
			for tmp != nil && tmp.val == tmpBack.val {
				tmp = tmp.right
				next = tmp
			}

			for pre != nil && pre.val == tmpBack.val {
				pre = pre.left
			}
			if next != nil {
				next.left = pre
			}
			pre.right = next
			fmt.Println("XXxxxxxx", pre, next)
			if next != nil {
				next.left = pre
			}
			tmp = tmpBack.down
		}
	}
	for tmp != nil {
		if tmp.right != nil {
			if target == tmp.right.val {
				del(tmp.right)
				return true
			} else if target < tmp.right.val {
				tmp = tmp.down
			} else {
				tmp = tmp.right
			}
		} else {
			tmp = tmp.down
		}
	}
	return false
}

/**

\* Your Skiplist object will be instantiated and called as such:

\* obj := Constructor();

\* param_1 := obj.Search(target);

\* obj.Add(num);

\* param_3 := obj.Erase(num);

*/
func objPrint(obj *Node) {

	pre := obj
	for obj != nil {
		pre = obj
		obj = obj.down
	}
	downToUp := pre
	fmt.Println("============= start format 1 =============")
	for downToUp != nil {
		preDownToUp := downToUp
		for downToUp != nil {
			fmt.Print(downToUp.val, "  ")
			downToUp = downToUp.up
		}
		downToUp = preDownToUp.right
		fmt.Println("\n-------")
	}
	fmt.Println("============= start format 2 =============")
	downToUp = pre
	for downToUp != nil {
		preDownToUp := downToUp
		for downToUp != nil {
			fmt.Print(downToUp.val, "  ")
			downToUp = downToUp.right
		}
		downToUp = preDownToUp.up
		fmt.Println("\n-------")
	}
	fmt.Println("============= end formats =============")
}

func printTime() {

	//timeBase := time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local)
	//fmt.Println(timeBase.Unix())
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
	objPrint(obj.head)

	fmt.Println(obj.Search(1))
	fmt.Println(obj.Search(200))
	fmt.Println(obj.Search(109))
	fmt.Println(obj.Search(107))
	fmt.Println(obj.Search(145))
	fmt.Println(obj.Search(123))
	fmt.Println(obj.Erase(100))
	objPrint(obj.head)
	fmt.Println(obj.Search(1999))
	fmt.Println(obj.Search(5))
	fmt.Println(obj.Search(123))
}

func runPrintTime() {
	printTime()
}
