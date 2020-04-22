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


## 动态规划

### 力扣：完全平方数
> https://leetcode-cn.com/problems/perfect-squares/
* 题解观点：一定要找到一个数组去保存过去的计算结果，避免重新计算消耗时间。包括斐波那契数列的计算。

### 力扣：最长上升子序列
> https://leetcode-cn.com/problems/longest-increasing-subsequence/
* 同样是建立一个dp数组