package disjoint_set_union

//讲解
//https://zhuanlan.zhihu.com/p/93647900/
//题目
//https://leetcode.cn/problems/largest-component-size-by-common-factor/submissions/

var father []int = make([]int, 1000)

func simpleFind(x int) int {
	if father[x] == x {
		return x
	} else {
		return simpleFind(father[x])
	}
}

func find(x int) int {
	if father[x] == x { //可以把p看成father，x访问到跟节点的标志就是father是自己
		return x
	} else {
		father[x] = find(father[x]) //father[x] = find(father[x]) //路径压缩，避免退化成长单链
	}
	return father[x]
}

func merge(a int, b int) {
	if find(a) == find(b) {
		return //
	}
	father[find(b)] = father[find(a)] // merge  a的祖先作为了合并的最终祖先。

}
