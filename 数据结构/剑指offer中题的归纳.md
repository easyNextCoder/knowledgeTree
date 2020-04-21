https://www.jianshu.com/p/53f6bf6f8d50

面试题57 - II. 和为s的连续正数序列
https://www.nowcoder.com/practice/c451a3fd84b64cb19485dad758a55ebe?tpId=13&tqId=11194&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking

力扣：求圆圈中最后剩余的数字
* 自己观点：约瑟夫环
* 题解观点：可以利用递推公式

力扣：面试题63：股票的最大利润
* 自己观点：穿上马甲的最大子列和
* 题解观点：还没看


力扣：求1+2+...+n；要求不能使用乘除法，for,while, if, else, switch,case等关键字及条件判断语句(A?B:C)
* 自己观点：不会
* 书上观点：使用构造函数，使用函数指针，使用虚函数，使用函数模板

力扣：不用加减乘除做加法
* a+b号可以分解为(a^b)+((a&b)<<1)
* 再令a = (a^b) b = ((a&b)<<1) 重复上面的步骤
* 直到b=0然后返回a,因为a+0 = a;

力扣：面试题66.构建乘积数组
* 自己观点：自己使用的是两层的乘法的循环
* 题解观点：构建两个数组用来存放前面0-i的乘积，和i-n的乘积

## 树

### 力扣：669修剪二叉搜索树
* 遇到树并都是唯一的一个操作：把树所有的节点按照顺序遍历一遍；而是可能会有其他更多的操作
> https://leetcode-cn.com/problems/trim-a-binary-search-tree/
### 力扣：判断一颗树是不是二叉树

>https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

注意：就是一个递归的判断每个节点是不是都是平衡二叉树的过程。每个节点都是相同的。

### 力扣：二叉树的深度

>https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/

### 力扣：二叉树的公共祖先
* 使用递归或者也可以使用链表的方法来进行判断
* **如何思考解决树的问题：**
	* （1）遍历树的函数中有travelleft() 和travelright()，请记住每个节点都会有这两个函数，所以如何在递归中处理就是在两个函数的**前 中 后**加入相应的代码
	* （2）无论是哪种遍历，最终都是将所有的节点无重复的遍历一遍，而并没有多余
	* （3）关于遍历的问题：
		* 1. 前序遍历（头 左体 右体）
		* 2. 中序遍历（左体 头 右体）（二叉搜索树，中序遍历完之后就是有序的数组）
		* 3. 后续遍历（左体 右体 头）


## 动态规划

### 力扣：完全平方数
> https://leetcode-cn.com/problems/perfect-squares/
* 题解观点：一定要找到一个数组去保存过去的计算结果，避免重新计算消耗时间。包括斐波那契数列的计算。

### 力扣：最长上升子序列
> https://leetcode-cn.com/problems/longest-increasing-subsequence/
* 同样是建立一个dp数组