package segment_tree

import "fmt"

//动态开点的方式，比较实用的
type Node struct {
	ls  int //左孩子在数组中的index
	rs  int //右孩子在数组的index
	add int //懒标记
	val int //代表当前节点有多少数
}

var count int = 0
var root []*Node

func dynamicUpdate(u, lc, rc, l, r, cnt int) {
	fmt.Println(u, lc, rc, l, r)
	if r < lc || l > rc {
		return
	}
	if lc >= l && rc <= r {
		root[u].add += cnt
		root[u].val += cnt * (rc - lc + 1)
		return //这里经常忘了return
	}
	create(u)
	pushDown(u, rc-lc+1)
	mid := (lc + rc) / 2
	dynamicUpdate(root[u].ls, lc, mid, l, r, cnt)
	dynamicUpdate(root[u].rs, mid+1, rc, l, r, cnt)
	pushUp(u)
}

func create(u int) {
	if root[u] == nil {
		root[u] = new(Node)
	}
	if root[u].ls == 0 {
		root[u].ls = count + 1
		count++
		root[root[u].ls] = new(Node)
	}

	if root[u].rs == 0 {
		root[u].rs = count + 1
		count++
		root[root[u].rs] = new(Node)
	}

}

func pushUp(u int) {
	root[u].val = root[root[u].ls].val + root[root[u].rs].val
}

func pushDown(u, len int) {
	rightChildIndex := root[u].rs
	leftChildIndex := root[u].ls
	root[rightChildIndex].add += root[u].add
	root[leftChildIndex].add += root[u].add
	root[rightChildIndex].val += root[u].add * (len / 2)
	//这里一定是root[u].add而不是root[rightChildIndex].add
	root[leftChildIndex].val += root[u].add * (len - len/2)
	root[u].add = 0
}

func dynamicQuery(u, lc, rc, l, r int) int {
	if r < lc || l > rc {
		return 0
	}
	if lc >= l && rc <= r {
		return root[u].val
	}
	create(u)
	pushDown(u, rc-lc+1)
	mid := (lc + rc) / 2
	return dynamicQuery(root[u].ls, lc, mid, l, r) + dynamicQuery(root[u].rs, mid+1, rc, l, r)
}

func runDynamic() {

	root = make([]*Node, 1000)
	count = 0
	dynamicUpdate(0, 1, 4, 1, 3, 1)
	res := dynamicQuery(0, 1, 4, 1, 4)
	fmt.Println(res, root[0], root[1], root[2], root)
}

//主席树的形式
func build(ori []int, tree []int, l, r, p int) {
	//l,r,p分别是左边起点右边起点和
	if l == r {
		tree[p] = ori[l]
		return
	}

	mid := l + (r-l)/2
	fmt.Println(l, r, p*2)
	build(ori, tree, l, mid, p*2)
	build(ori, tree, mid+1, r, p*2+1)
	tree[p] = tree[2*p] + tree[2*p+1]
}

func update(lazy []int, tree []int, treeLeft, treeRight, cnt, start, end, pos int) {
	if start > treeRight || end < treeLeft {
		return
	}
	// [treeLeft,treeRight] 为修改区间,count 为被修改的元素的变化量,[start,end] 为当前节点包含的区间,pos
	// 为当前节点的编号
	if treeLeft <= start && end <= treeRight {
		//当前节点对应的区间包含在目标区间中，则需要全部更改
		tree[pos] += (end - start + 1) * cnt
		//先把所有的增量增给父亲节点
		lazy[pos] += cnt
		//同时用懒标记数组记录下，每个儿子应该分配的值
		return
	}
	m := (start + end) / 2
	lazy[pos*2+1] += lazy[pos]
	//标记向下传递
	lazy[pos*2+2] += lazy[pos]

	tree[pos*2+1] += lazy[pos] * (m - start + 1)
	//往下更新一层
	tree[pos*2+2] += lazy[pos] * (end - m)
	lazy[pos] = 0
	//清除标记
	update(lazy, tree, treeLeft, treeRight, cnt, start, m, pos*2+1)
	//递归的往下寻找
	update(lazy, tree, treeLeft, treeRight, cnt, m+1, end, pos*2+2)
	tree[pos] = tree[pos*2+1] + tree[pos*2+2]
	//根据子节点更新当前节点的值
}

//update2是自己写的
func update2(tree, lazy []int, p, treeLeft, treeRight, start, end, cnt int) {
	//与update中参数的定义完全相反
	if treeLeft > end || treeRight < start {
		return
	}

	if start <= treeLeft && treeRight <= end {
		tree[p] += (treeRight - treeLeft + 1) * cnt
		lazy[p] += cnt
		return
	}
	mid := (treeLeft + treeRight) / 2
	//lazy更新
	lazy[p*2+1] += lazy[p]
	lazy[p*2+2] += lazy[p]

	//tree更新
	tree[p*2+1] += (mid - treeLeft + 1) * lazy[p]
	tree[p*2+2] += (treeRight - mid) * lazy[p]
	lazy[p] = 0
	update2(tree, lazy, p*2+1, treeLeft, mid, start, end, cnt)    //left
	update2(tree, lazy, p*2+2, mid+1, treeRight, start, end, cnt) //right
	tree[p] = tree[p*2+2] + tree[p*2+1]
}

func query(lazy []int, tree []int, p, treeL, treeR, start, end int) int {
	if treeL > end || treeR < start {
		return 0
	}
	if treeL >= start && treeR <= end {
		return tree[p]
	}
	mid := (treeL + treeR) / 2
	tree[p] = query(lazy, tree, p*2+1, treeL, mid, start, end) + query(lazy, tree, p*2+2, mid+1, treeR, start, end)
	return tree[p]
}

func runBuild() {
	ori := []int{1, 2, 3, 4, 5}
	tree := make([]int, 1<<len(ori)+1)
	lazy := make([]int, 1<<len(ori)+1)
	//创建树
	build(ori, tree, 0, 4, 1)
	queryRes := query(lazy, tree, 0, 0, 4, 1, 3)
	fmt.Println(lazy, tree)
	fmt.Println(queryRes)
	//更新树
	update(lazy, tree, 1, 3, 4, 0, 4, 0)
	//update2(tree, lazy, 0, 0, 4, 1, 3, 4)
	queryRes2 := query(lazy, tree, 0, 0, 4, 1, 3)
	fmt.Println(lazy, tree)
	fmt.Println(queryRes2)

}
