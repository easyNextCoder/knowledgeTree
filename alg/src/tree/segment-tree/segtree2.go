package segment_tree

import "fmt"

type Node2 struct {
	leftChildIdx  int
	rightChildIdx int
	add           int
	val           int
}

var root2 []*Node2
var count2 int = 0

func create2(uc int) {
	//if root2[uc] != nil {
	//	return
	//}
	if root2[uc] == nil {
		root2[uc] = &Node2{
			leftChildIdx:  count2 + 1,
			rightChildIdx: count2 + 2,
			add:           0,
			val:           0,
		}
		count2 += 2
		leftChild := root2[uc].leftChildIdx
		rightChild := root2[uc].rightChildIdx
		root2[leftChild] = new(Node2)
		root2[rightChild] = new(Node2)
		fmt.Println("xxx")
	}

	if root2[uc].leftChildIdx == 0 {
		root2[uc].leftChildIdx = count2 + 1
		fmt.Println("xxxx1", uc, root2[uc].leftChildIdx)
		count2++
		root2[root2[uc].leftChildIdx] = new(Node2)
		fmt.Println("xxxx2", uc, root2[uc].leftChildIdx)
	}

	if root2[uc].rightChildIdx == 0 {
		root2[uc].rightChildIdx = count2 + 1
		fmt.Println("xxxxx1", uc, root2[uc].rightChildIdx)
		count2++
		root2[root2[uc].rightChildIdx] = new(Node2)
		fmt.Println("xxxxx2", uc, root2[uc].rightChildIdx)
	}
	fmt.Println("yy", root2[uc].leftChildIdx, root2[uc].rightChildIdx)
}
func pushD(uc, cnt, len int) {
	leftChild := root2[uc].leftChildIdx
	rightChild := root2[uc].rightChildIdx

	root2[leftChild].val += len / 2 * root2[uc].add
	root2[leftChild].add += root2[uc].add
	root2[rightChild].val += (len - len/2) * root2[uc].add
	root2[rightChild].add += root2[uc].add
	root2[uc].add = 0
}

func pushU(uc int) {
	root2[uc].val = root2[root2[uc].leftChildIdx].val + root2[root2[uc].rightChildIdx].val
}

func updateSeg(uc, lc, rc, l, r, cnt int) {
	fmt.Println("a root2[uc]:", uc, lc, rc, l, r, root2[uc])
	if l > rc || r < lc {
		fmt.Println("b root2[uc]", uc, lc, rc, l, r, root2[uc])
		return
	}

	if l <= lc && rc <= r {
		root2[uc].add += cnt
		root2[uc].val += (rc - lc + 1) * cnt
		fmt.Println("c root2[c]", uc, lc, rc, l, r, root2[uc])
		return
	}
	fmt.Println("d root2[uc]:", uc, lc, rc, l, r, root2[uc])
	create2(uc)
	pushD(uc, cnt, rc-lc+1)
	midBorder := (lc + rc) / 2
	fmt.Println("dd root2[uc]:", uc, root2[uc].leftChildIdx, root2[uc].rightChildIdx)
	updateSeg(root2[uc].rightChildIdx, midBorder+1, rc, l, r, cnt)
	updateSeg(root2[uc].leftChildIdx, lc, midBorder, l, r, cnt)

	pushU(uc)
	fmt.Println("e root2[uc]:", uc, lc, rc, l, r, root2[uc])

}
func dynamicQuery2(u, lc, rc, l, r int) int {
	if r < lc || l > rc {
		return 0
	}
	if lc >= l && rc <= r {
		return root2[u].val
	}
	create2(u)
	pushD(u, root2[u].add, rc-lc+1)
	mid := (lc + rc) / 2
	return dynamicQuery2(root2[u].leftChildIdx, lc, mid, l, r) + dynamicQuery2(root2[u].rightChildIdx, mid+1, rc, l, r)
}
func runDynamic2() {
	root2 = make([]*Node2, 1000)
	count = 0
	updateSeg(0, 1, 4, 1, 3, 1)
	res := 0 //dynamicQuery2(0, 1, 4, 1, 3)
	fmt.Println(res, root2[0], root2[1], root2[2], root2)

}
