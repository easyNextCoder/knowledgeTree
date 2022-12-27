package segtree_1

import "fmt"

type Node struct {
	left  int
	right int
	add   int
	val   int
}

func (s *Node) String() string {
	return fmt.Sprintf("(left:%d,right:%d,add:%d,val:%d)", s.left, s.right, s.add, s.val)
}

var root []*Node
var count int = 1

func create(u int) {
	if root[u] == nil {
		root[u] = new(Node)
	}

	if root[u].left == 0 {
		root[u].left = count
		root[count] = new(Node)
		count++
	}

	if root[u].right == 0 {
		root[u].right = count
		root[count] = new(Node)
		count++
	}
}

func pushd(u int, len int) {
	left := root[u].left
	right := root[u].right

	root[left].add += root[u].add
	root[right].add += root[u].add
	root[left].val += root[u].add * len / 2
	root[right].val += root[u].add * (len - len/2)
	root[u].val = 0
}
a := []int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}
b := []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}
func pushu(u int) {
	root[u].val = root[root[u].left].val + root[root[u].right].val
}

func update(u, lc, rc, l, r, cnt int) {
	if l > rc || r < lc {
		return
	}

	if l <= lc && r >= rc {
		root[u].add += cnt
		root[u].val += (rc - lc + 1) * cnt
		return
	}

	create(u)
	pushd(u, rc-lc+1)
	mid := (lc + rc) / 2
	update(root[u].left, lc, mid, l, r, cnt)
	update(root[u].right, mid+1, rc, l, r, cnt)
	pushu(u)

}

func run() {
	root = make([]*Node, 1000)
	update(0, 1, 4, 1, 3, 1)
	fmt.Println(root)
}
