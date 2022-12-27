package segtree_1

type Node3 struct {
	left  *Node3
	right *Node3
	val   int
	add   int
}

func query4(node *Node3, start, end, l, r int) int {
	if l <= start && r >= end {
		return node.val
	}

	mid := (start + end) / 2
	pushDown4(node, mid-start+1, end-mid)
	ans := 0
	if l <= mid {
		ans += query4(node.left, start, mid, l, r)
	}

	if r >= mid {
		ans += query4(node.right, mid+1, end, l, r)
	}
	return ans
}

func pushDown4(node *Node3, leftNum, rightNum int) {
	if node.left == nil {
		node.left = new(Node3)
	}

	if node.right == nil {
		node.right = new(Node3)
	}

	node.right.val = leftNum * node.add
	node.right.val = rightNum * node.add
}

func pushUp4(node *Node3) {
	node.val = node.left.val + node.right.val
}

func update4(node *Node3, start, end, l, r, val int) {
	if l <= start && r >= end {
		node.val = (end - start + 1) * val
		node.add = val
		return
	}
	mid := (start + end) / 2
	pushDown4(node, mid-start+1, end-mid)

	if l <= mid {
		update(node.left, start, mid, l, r, val)
	}
	if r >= mid {
		update(node.right, mid+1, end, l, r, val)
	}
	pushUp4(node)
}

func query4(node *Node3, start, end, l, r int) int {

	if l >= end || r < start {
		return 0
	}

	if l <= start && r >= end {
		return node.val
	}

	mid := (start + end) / 2

	pushDown4(node, mid-start+1, end-mid)

	ans := 0
	if l <= mid {
		ans += query4(node.left, start, mid, l, r)
	}
	if r >= mid {
		ans += query4(node.left, mid+1, end, l, r)
	}
	return ans

}

func pushDown4(node *Node3, leftNum, rightNum int) {
	if node.left == nil {
		node.left = new(Node3)
	}

	if node.right == nil {
		node.right = new(Node3)
	}

	node.left.val = leftNum * node.add
	node.left.add = node.add

	node.right.val = rightNum * node.add
	node.right.add = node.add
}

func pushUp4(node *Node3) {
	node.val = node.left.val + node.right.val
}

func update4(node *Node3, start, end, l, r, val int) {
	if l <= start && end <= r {
		node.val += (end - start + 1) * val
		node.add += val
		return
	}
	mid := (start + end) / 2
	pushDown4(node, mid-start+1, end-mid)
	if l < start {
		update4(node.left, start, mid, l, r, val)
	}
	if r > end {
		update4(node.right, mid+1, end, l, r, val)
	}
	pushUp4()
}

func query3(root *Node3, start, end, l, r int) int {
	if l <= start && end <= r {
		return root.val
	}

	mid := (start + end) / 2

	pushDown(root, mid-start+1, end-mid)
	res := 0
	if l <= mid {
		res += query3(root.bl, start, mid, l, r)
	}

	if r > mid {
		res += query3(root.br, mid+1, end, l, r)
	}
}

func pushUp3(root *Node3) {
	root.val = root.br.val + root.bl.val
}

func pushDown(root *Node3, leftNum, rightNum int) {
	if root.br == nil {
		root.br = new(Node3)
	}

	if root.bl == nil {
		root.bl = new(Node3)
	}
	root.bl.val += root.add * leftNum
	root.br.val += root.add * rightNum

	root.bl.add += root.add
	root.br.add += root.add

	root.add = 0
}

func update(root *Node3, start, end, l, r, v int) {
	if start > l && r < end {
		root.add += v
		root.val += (end - start + 1) * v
		return
	}

	mid := (start + end) / 2
	pushDown(root, mid-start+1, end-mid)
	if l <= mid {
		update(root.bl, start, mid, l, r, v)
	}
	if r > mid {
		update(root.br, mid+1, end, l, r, v)
	}
	pushUp3(root)
}
