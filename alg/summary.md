

# 资源

[网址](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)

# 数组

* 区间数组

  * 判断是否相交

    (start0, end0) (start1, end1)

    start0 > end1 || end0 <start1

  * 找到相交区间
    [max(start0, start1), min(end0, end1)]

## 题目

* * 区间是否相交判断方法
    http://t.zoukankan.com/liuwt365-p-7222549.html


# 排序

[排序总结](https://babyvector.github.io/blog/MD-sort-algorithm-tutorial.html)

## 拓扑排序

在AOV网中，若不存在回路，则所有活动可排列成一个线性序列，使得每个活动的所有前驱活动都排在该活动的前面，我们把此序列叫做拓扑序列(Topological order)，由AOV网构造拓扑序列的过程叫做拓扑排序(Topological sort)。AOV网的拓扑序列不是唯一的，满足上述定义的任一线性序列都称作它的拓扑序列。

# 多指针

## 极值指针

#### [581. 最短无序连续子数组](https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/)

## 双指针

* 双指针删除重复元素

  #### [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

## 快慢指针

# 二分法

二分查找的核心是找到具有单调性的正确值



二分法一共三种写法

* 按照[left, right]进行处理

  ```java
  while(left<=right){
  	int mid = left+(right-left)/2;
  	if(arr[mid] == k){
  		return mid;	
  	}else if(arr[mid] > k){
  		right = mid-1;
  	}else{
  		left = mid+1;
  	}
  }
  ```

  能够对区间内的所有元素进行遍历

* 按照[left, right)进行处理

  ```java
  while(left<right){//等号没有意义
  	int mid = left+(right-left)/2;
  	if(arr[mid] == k){
  		return mid;
  	}else if(arr[mid]>k){
  		right = mid;
  	}else{
  		left = mid+1;
  	}
  }
  ```

  这里对right是无法遍历到的
  以上两种方法的[链接](https://blog.csdn.net/gongkeguo/article/details/123255001?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-123255001-blog-112527549.pc_relevant_antiscanv2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-123255001-blog-112527549.pc_relevant_antiscanv2&utm_relevant_index=5)

* 划分红蓝区间进行处理

  ```java
  while(left+1<right){//保证一定能将整个数组划分位红蓝两个部分
  	int mid = left+(right-left)/2;//取下边界值
  	if(arr[mid] == k){
  		return mid;
  	}else if(k<arr[mid]){//目标值k在左边子数组中，这样写更容易理解
  		right = mid;
  	}else{
  		left = mid;
  	}
  }
  ```

  [原文链接](https://www.bilibili.com/video/BV1d54y1q7k7?spm_id_from=333.880.my_history.page.click)

* 相关题目

  https://leetcode.cn/problems/nZZqjQ/submissions/

  https://leetcode.cn/problems/jJ0w9p/

  #### [29. 两数相除](https://leetcode.cn/problems/divide-two-integers/)

## 二分法的工程实用

* lower_bound()

  ```c++
  int myLowerBound(vector<int> &data, int k)
  {
  	int start = 0;
  	int last = data.size();
  	while (start < last)
  	{
  		int mid = (start + last) / 2;
  		if (data[mid] >= k)
  		{
  			last = mid;
  		}
  		else
  		{
  			start = mid + 1;
  		}
  	}
  	return start;
  }
  ```

  

* Upper_bound()

  ```c++
  int myUpperBound(vector<int> &data, int k)
  {
  	int start = 0;
  	int last = data.size();
  	while (start < last)
  	{
  		int mid = (start + last) / 2;
  		if (data[mid] <= k)
  		{
  			start = mid + 1;
  		}
  		else
  		{
  			last = mid;
  		}
  	}
  	return start;
  }
  ```

* 自己的测试代码

  ```go
  arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 7, 8, 10}
  	upper_bound := func() int {
  		target := 5
  		l, r := 0, len(arr)
  		for l < r {
  			mid := l + (r-l)/2
  			if target >= arr[mid] {
  				l = mid + 1
  			} else {
  				r = mid
  			}
  		}
  		return l
  	}
  	ret := upper_bound()
  	fmt.Println(ret)
  
  ```

  

# 位运算

## 题目

* 4字节int类型的最大值和最小值

* 正码、反码、补码（最大值，最小值的二进制）、4字节int越界问题
  [231. 2 的幂](https://leetcode.cn/problems/power-of-two/)

  #### [29. 两数相除](https://leetcode.cn/problems/divide-two-integers/)

# 树

## 基本知识

高度为h的满二叉树的节点个数为：2^(h)-1个，叶子节点个数为2^(h-1)

## AVL树

在计算机科学中，AVL树是最早被发明的自平衡二叉查找树。在AVL树中，任一节点对应的两棵子树的最大高度差为1，因此它也被称为高度平衡树。查找、插入和删除在平均和最坏情况下的时间复杂度都是{\displaystyle O(\log {n})}O(\log{n})。增加和删除元素的操作则可能需要借由一次或多次树旋转，以实现树的重新平衡。

算法        平均        最差
空间        O(n)        O(n)
搜索        O(log n)    O(log n)
插入        O(log n)    O(log n)
删除        O(log n)    O(log n)

平衡二叉树类型    平衡度    调整频率   适用场景
AVL树           高       高       查询多，增/删少
红黑树            低       低       增/删频繁



## 前缀树

## 线段树

### [线段树讲解](https://www.cnblogs.com/RioTian/p/13409694.html)

* 主席树形式
  线段树（Segment Tree）是一种二叉树形数据结构，每个节点p的左右子节点编号分别为2p和2p+1，假如节点p存储区间[a,b]的和，设mid=[(l+r)/2], 那么两个子节点分别存储[l, mid]和[mid+1, r]的和。

* 动态开点形式
  对于大范围的范围查询一般使用这种方式
* 题目
  [731. 我的日程安排表 II](https://leetcode.cn/problems/my-calendar-ii/)

## B树

* B树：即普通的二叉搜索树

* B-树：是一种多路搜索树（并不是二叉的）
      B-树的特征：
          1.关键字集合分布在整棵树中；
          2.任何一个关键字出现且只出现在一个节点中；
          3.搜索有可能在非叶子节点结束
          4.其搜索性能等价于在关键字全集内做一次二分查找
          5.自动层次控制
          搜索性能：O(log2 N)

* B+树：是B-树的变体，也是一种多路搜索树
      B+树的定义：
          1.其定义基本于B-树相同，除了
          2.非叶子节点的子树指针与关键字个数相同；
          3.非叶子节点的子树指针P[i],指向关键字值属于[K[i], K[i+1]]的子树（B-树是开区间）
          4.为所有叶子节点增加一个链指针
          5.所有关键字都在叶子节点出现
      B+树的特征：
          1.所有关键字都出现在叶子节点的链表中（稠密索引），且链表中的关键字恰好是有序的
          2.不可能在非叶子节点命中
          3.非叶子节点相当于是叶子节点的索引（稀疏索引），叶子节点相当于是存储（关键字）数据的数据层
          4.更适合文件索引系统

* 总结：
      B树：二叉树，每个节点只存储一个关键字，等于则命中，小于走左节点，大于走右节点
      B-树：多路搜索树，每个节点存储M/2到M个关键字，非叶子节点存储指向关键字范围的子节点，所有关键字在整棵树中出现，且只出现一次，非叶子节点可以命中
      B+树：在B-树的基础上，为叶子节点增加链表指针，所有关键字都在叶子节点中出现，非叶子节点作为叶子节点的索引；B+树总是到叶子节点才命中

## 红黑树

RB-tree不仅是一个二叉搜索树，而且必须满足以下规则：
1. 每个节点不是红色就是黑色
2. 根节点为黑色
3. 如果节点为红，其子节点必须为黑。
4. 任一节点至NULL（树尾端）的任何路径，所含的黑节点数必须相同

STL中红黑树属于双向迭代器，但是不具备随机定位能力

## 题目

#### [230. 二叉搜索树中第K小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)

这个是有两种方法的，两种方法自己都做了

#### [406. 根据身高重建队列](https://leetcode.cn/problems/queue-reconstruction-by-height/)

线段树 or 数组排序



# 堆

## 大小顶堆

* 堆排序

# 递归

## 回溯

## 分治

## 题目

# 动态规划

## 能用递归写的，就大部分能改成动态规划

[871. 最低加油次数](https://leetcode.cn/problems/minimum-number-of-refueling-stops/)

```go
func min(a, b int)int{
    if a<b{
        return a
    }
    return b
}
func work(stations [][]int, index int, fuel int, meilHaveRun int, totalMeil int, cnt int)int{
    if meilHaveRun+fuel >= totalMeil {
        return cnt
    }
    if index >= len(stations){
        return 10000
    }
    if fuel + meilHaveRun >= stations[index][0]{
        res1 := work(stations, index+1, fuel, meilHaveRun, totalMeil, cnt)//不从这个加油站里加油
        
        res2 := work(stations, index+1, fuel-(stations[index][0]-meilHaveRun)+stations[index][1], stations[index][0], totalMeil, cnt+1)//从这个加油站里加油
        
        return min(res1, res2)
    }else{
        return 10000
    }
}

func max(a, b int) int{
    if a>b{
        return a
    }
    return b
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
    //dp[i][j] 表示对于 [0, i - 1] 范围内的加油站，最多加 j 次油能够到达的最远距离。
    dp := make([][]int, len(stations)+1)
    for k, _ := range dp{
        dp[k] = make([]int, len(stations)+1)
    }
    dp[0][0] = startFuel
    finalRes := 100000
    xLen := len(stations)
    for i:=1; i<=xLen; i++{
        dp[i][0] = startFuel
        for j:=1; j<=xLen; j++{
            //不从这个加油站里加油
            if dp[i-1][j] >= stations[i-1][0]{
                dp[i][j] = dp[i-1][j]
            }
            //从这个加油站加油
            if dp[i-1][j-1] >= stations[i-1][0]{
                res1 := dp[i-1][j-1] + stations[i-1][1]
                dp[i][j] =max(dp[i][j], res1)
            }
        }
    }
    // fmt.Println(dp)
    for _, v := range dp{
        for j, vv:=range v{
            if vv >= target{
                finalRes = min(finalRes, j)
            }   
        }
    }
    if finalRes == 100000{
        finalRes = -1
    }
    return finalRes
}
```

我觉得吧，官方题解的动态规划写得并不好，没有从最基本的状态转移方程开始推导，直接把二维的状态压缩成一维的，这对很多朋友都不够友好。

贴一个二维状态定义的动态规划。

状态定义： `dp[i][j]` 表示对于 `[0, i - 1]` 范围内的加油站，最多加 `j` 次油能够到达的最远距离。

初始化条件： （1）当 `j == 0` 时，此时有 `dp[i][0] = startFuel`。 （2）当 `i < j` 时，这种情况其实是不存在的，因为总共只有 `i` 个加油站却需要加 `j` 次油，设成 0 代表不存在。

状态转移方程： （1）我们可以选择在第 `i - 1` 个加油站加油，此时 `dp[i][j] = dp[i - 1][j - 1] + stations[i - 1][1]`。当然能进行这个状态转移的条件是`dp[i - 1][j - 1] >= stations[i - 1][0]`，因为必须要到达第 `i - 1` 个加油站才能选择在第 `i - 1` 个加油站加油与否。 （2）我们也可以选择不在第 `i - 1` 个加油站加油，此时的状态转移方程更简单，是`dp[i][j] = dp[i - 1][j]`。 取上述 `2` 种情况的最大值就是我们的 `dp[i][j]`。

为了降低大家的理解难度，用记忆化搜索的方式实现了上述动态规划解法，其中的 `helper(i, j)` 函数就对应上文分析中的状态定义 `dp[i][j]`。

时间复杂度和空间复杂度均是 `O(n ^ 2)`。[官方题解第一个评论](https://leetcode.cn/problems/minimum-number-of-refueling-stops/solution/zui-di-jia-you-ci-shu-by-leetcode-soluti-nmga/)

## 选择问题

### 01背包模型

* 简单背包
* [完全背包问题](https://www.acwing.com/problem/content/3/)
* [多重背包问题](https://www.acwing.com/problem/content/4/)
* [二维费用背包问题](https://www.acwing.com/problem/content/8/)
  * [服务部署](https://www.nowcoder.com/login?callBack=%2Fprofile%2F106427925%2FcodeBookDetail%3FsubmissionId%3D78195067%26headNav%3Dwww)
* [背包问题求方案数](https://www.acwing.com/problem/content/description/11/)
  * 题目：[数字组合](https://www.acwing.com/problem/content/280/)
  * 题目：[自然数拆分](https://www.acwing.com/problem/content/281/)

### 多约束背包问题

* [牛客网-运矿石](https://www.nowcoder.com/questionTerminal/b58f922cc924478fa1e2dca2cc4f4eb7)

## 区间dp问题

* 基本问题
  * 力扣96-[不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/)
  * acwing-[石子合并](https://www.acwing.com/problem/content/284/)
* 带有限制条件的区间dp问题
  * 牛客网-[vivo一维消消乐](https://www.nowcoder.com/profile/106427925/test/33760813/637397?&headNav=www)
    * [546. 移除盒子](https://leetcode.cn/problems/remove-boxes/)
  * 力扣1000-[合并石头的最低成本](https://leetcode.cn/problems/minimum-cost-to-merge-stones/submissions/)
  * 力扣312-[戳气球](https://leetcode.cn/problems/burst-balloons/)
    其实动态规划的规程就是一个分解问题成独立子问题然后求解的过程，难点一般都是难在如何将所有的过程都划分为独立的子过程。对于这道题，如果正向思维来看戳爆一个气球之后它并不是一个子过程逆向思考可以把他们转化为一个子过程。

### 记忆搜索问题

* 力扣1139-[最大的以 1 为边界的正方形](https://leetcode.cn/problems/largest-1-bordered-square/)

### 序列问题

* 最佳买卖股票时机含冷冻期
  [309. 最佳买卖股票时机含冷冻期](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

* 打家劫舍问题
  [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)

* 打家劫舍变种（也可以看成区间dp问题）
  [740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/)

* 最长上升子序列

  [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

  问题解析1：(暴力不可行)

  可以使用递归或者位操作的方式枚举出所有可能的序列：复杂度：O(2^N)再找出上升的O(N2^N)
  问题解析2：（动态规划可行）

  dp[i]表示以i结尾的最长上升子序列

* 最长上升子序列变种：

  力扣：354-[俄罗斯套娃信封问题](https://leetcode.cn/problems/russian-doll-envelopes/)

### 序列极值问题

求序列的子序列的最大和

* 力扣523-[连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
  一共有两种解法

  1.复杂度为O(n^2)

  2.复杂度为O(N)使用hashset对运算进行了加速

* 力扣1191-[K次串联后最大子数组之和](https://leetcode.cn/problems/k-concatenation-maximum-sum/)

  

### 双字符串动态规划

* 最长公共子序列

  * 最长公共子序列的定义

    在数学中，某个序列的子序列是从最初序列通过去除某些元素但不破坏余下元素的相对位置（在前或在后）而形成的新序列。

    例子：[a,b,c,d,e]的子序列可以是[a,b,c][a,c,e][b,e]但是不可以是[c,a,e]

  * 解决方案

    ```
    我们使用一个二维数组 dpdp， dp[i][j]dp[i][j] 表示 s1s1 前 ii 个字符和 s2s2 前 jj 个字符中最长公共子序列。我们逐行填充 dpdp 数组。
    
    对于每一个 dp[i][j]dp[i][j]，我们有 2 种选择：
    
    字符 s1[i-1]s1[i−1] 和 s2[j-1]s2[j−1] 匹配，那么 dp[i][j]dp[i][j] 会比两个字符串分别考虑到前 i-1i−1 个字符 和 j-1j−1 个字符的公共子序列长度多 1 。所以 dp[i][j]dp[i][j] 被更新为 dp[i][j] = dp[i-1][j-1] + 1dp[i][j]=dp[i−1][j−1]+1。注意到 dp[i-1][j-1]dp[i−1][j−1] 已经被求解过了，所以可以直接使用。
    
    字符 s1[i-1]s1[i−1] 和 s2[j-1]s2[j−1] 不匹配，这种情况下我们不能直接增加已匹配子序列的长度，但我们可以将之前已经求解过的最长公共子序列的长度作为当前最长公共子序列的长度。但是我们应该选择哪一个呢？事实上此时我们有 2 种选择。我们可以删除 s1s1 或者 s2s2 的最后一个字符然后将对应的 dpdp 数组的值作比较，也就是取 dp[i-1][j]dp[i−1][j] 和 dp[i][j-1]dp[i][j−1] 的较大值。
    
    dp[m][n]就是两个字符串的最长公共子序列
    
    动态规划的转移公式为：
    if(s1[i-1] == s2[j-1])
    	dp[i][j] = dp[i-1][j-1]+1;
    else
    	dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
    ```

  * 例题讲解

    * 力扣72-[编辑距离](https://leetcode.cn/problems/edit-distance/)
    * 力扣583-[两个字符串的删除操作](https://leetcode.cn/problems/delete-operation-for-two-strings/)
    * 力扣712-[两个字符串的最小ASCII删除和](https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/)

## 其他

### 状态机的模型

### 状态压缩dp的模型

* 力扣1434-[每个人戴不同帽子的方案数](https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/)

### 树形dp的模型

### 单调队列优化？

### 斜队列优化？



# 图

* 基本定义

  入度：有向图中箭头指向自己的线的个数

  出度：有向图中箭头从自己身上出的线的个数

  拓扑排序

  * 定义

    对一个[有向无环图](https://baike.baidu.com/item/有向无环图/10972513)(Directed Acyclic Graph简称DAG)G进行拓扑排序，是将G中所有顶点排成一个线性序列，使得图中任意一对顶点u和v，若边<u,v>∈E(G)，则u在线性序列中出现在v之前。

  * 过程

    先将入度为0的点加入到队列s（如果遇到多个节点入度为0则说明拓扑排序不唯一）
    将s指向的节点的入度减1，将该元素append到out数组中

    重复以上步骤，若输出的顶点个数小于网格中顶点的个数，则说明图中存在环

  

* 基本表示方法

  ``g[][]``表示整个图的结构

  ``inDeg[]``记录图的入度

  ``outDeg[]``记录图的出度

## 题目

#### [剑指 Offer II 115. 重建序列](https://leetcode.cn/problems/ur2n8P/)

# 链表

## 哨兵

## 题目

[剑指 Offer II 029. 排序的循环链表](https://leetcode.cn/problems/4ueAj6/)

本题如果逻辑梳理清晰，可以做到一次遍历插入，具体官方答案中有

# 贪心

## 题目





# 状态机

#### [剑指 Offer 20. 表示数值的字符串](https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

**有时间去学习编译原理**

**有时间把自己的jos做完**

# 字符串

## kmp算法





# 习题

### Leetcode题目

#### [146. LRU 缓存](https://leetcode.cn/problems/lru-cache/)

双向链表+hashmap

#### [169. 多数元素](https://leetcode.cn/problems/majority-element/)

摩尔投票大乱斗

#### [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/)

深度或者广度优先搜索

时间复杂度O(MN)

#### [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)

前缀树

#### [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)

单调队列，deque，O(1)在任意时刻取出数组的最大值

ArrayList<Integer> -> int[]  <=> arr.stream().mapToInt(Integer::intValue).toArray();

#### [215. 数组中的第K个最大元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/)

默认的堆都是小顶堆

#### [240. 搜索二维矩阵 II](https://leetcode.cn/problems/search-a-2d-matrix-ii/)

二分查找和分治

#### [283. 移动零](https://leetcode.cn/problems/move-zeroes/)

双指针

#### [287. 寻找重复数](https://leetcode.cn/problems/find-the-duplicate-number/)

数组

#### [297. 二叉树的序列化与反序列化](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/)

字符串处理、树的层序遍历

#### [338. 比特位计数](https://leetcode.cn/problems/counting-bits/)

位运算的优先符号，低于逻辑运算的符号

#### [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)

堆，大顶堆和小顶堆，O(nlogk)前K大元素问题

#### [394. 字符串解码](https://leetcode.cn/problems/decode-string/)

栈

#### [990. 等式方程的可满足性](https://leetcode.cn/problems/satisfiability-of-equality-equations/)

#### [399. 除法求值](https://leetcode.cn/problems/evaluate-division/)（带权重的并查集合）

并查集

* 并查集（不相交集合）用于处理**动态连通性**问题，最典型的应用是求几个最小生成树的Kruskal算法
* 并查集支持（1）查询find（2）合并union两个操作
* 并查集只回答两个节点是不是在一个联通分量中（也就是所谓的连通性问题），并不回答路径问题
* 如果一个问题具有传递性，可以考虑使用并查集
* 并查集最常见的一种设计思想是：把同在一个连通分量中的节点组织成一个树形结构
