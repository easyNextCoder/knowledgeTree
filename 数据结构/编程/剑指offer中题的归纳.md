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



## 以下是待归纳的项目：

leetcode 牛客网 
350题：两个数组的交集2
问题：其中使用了快排算法不会
https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/


leetcode 牛客网
10题： I. 斐波那契数列
问题：斐波那契数的求解如果使用递归来做，小数超出时间限制；大数会暴栈
解决方案：使用java中的BigInteger在java.math.*包

leetcode 牛客网
12题： 矩阵中的路径
问题：考察使用的是DFS，但是自己的代码比较复杂
解决方案：以后重新写并简化代码；自己重新使用DFS代码再重新写一下
面试题13. 机器人的运动范围
是广度优先算法的考点

leetcode 整数拆分（最终自己没有总结出来）
要用到相关数学证明：结论就是让拆分的3尽量多
证明：https://leetcode-cn.com/problems/integer-break/solution/zheng-shu-chai-fen-shu-xue-fang-fa-han-wan-zheng-t/
面试题14- II. 剪绳子 II
这次使用到了BigInteger中的pow->函数原型pow(int n)
面试题15. 二进制中1的个数
原码-》取反-》反码-》加一-》补码
public class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int toDivde = n;
        int neg = 0;
        int count = 0;
        if(toDivde < 0){
            neg = 1;
            toDivde = ~toDivde;
        }
        while(toDivde != 0){
            if(toDivde %2 == 1){
                count++;
                toDivde/=2;
            }else{
                toDivde/=2;
            }
        }
        if(neg == 1)
            return 32-count;
        else
            return count;
    }
}

面试题16. 数值的整数次方（这个不会）求：1.000001^123456    
问题：pow次数过大超时，小数位数过多精度不够

自己的算法：
int cutoff = 1;
            if(b > 50){
                result = x;
                
                while(b-cutoff > 0 && b-2*cutoff>0){
                    result*=result;
                    cutoff*=2;
                }
                b-=cutoff;
                while(b-->0){
                    result*=x;
                }
            }else{
                result = 1;
                while(b-->0){
                    result*=x;
                }
            }
遇到超级大的数字的指数运算的时候，仍然会超时；好方法见网页上

最终自己的代码改成如下形式完美通过：
                double finalResult = 1;
                result = x;
                while(b>1000){
                      
                    while(b-cutoff > 0 && b-2*cutoff>0 && 2*cutoff>cutoff){//乘以2147483647会越界(越界之后仍然满足b-2*cutoff>0)
                        result*=result;
                        cutoff*=2;
                    }
                    b-=cutoff;
                    finalResult*=result;
                    cutoff = 1;    
                    result = x;
                }
                while(b-->0){
                    finalResult*=x;
                }
                result = finalResult;
只有b-cutoff>0->计算错误
只有b-cutoff>0 && b- 2*cutoff>0 ->计算超时错误，因为2*cutoff对于大整数而言会产生越界！


本题中用到的知识总结：
float类型：1符号位 8 23
double类型：1符号位 11 52
int占4个字节
long占8个字节 所以只用int无法表示与int负的最大值对应的正值，只能用long类型来表示


普通int型数据转化为二进制位：不停的除
普通的小于0的浮点型数据转化成二进制：不停的乘以2，得到的结果大于1则取1，否则取0；直到乘以2之后结果为0


面试题17. 打印从1到最大的n位数

        while(n-- >0){
            end*=10;
        }//注意n--和--n的区别

        若n = 1；n--是一定会运行一次，而--n则一次都不会运行

面试题20. 表示数值的字符串（这题自己写了一天，太久了）
注意测试样本：
"46.e3"
"-.8"
"-1E-16"
"0"
"+100"
"5e2"
"-123"
"3.1416"
"0123"
"-123"
"3.1416"
"0123"
"-1E-16"
"12e"
"1a3.14"
"1.2.3"
"+-5"
"12e+5.4"
"-."
true
true
false
true
true
true
true
true
true
true
true
true
false
false
false
false
false
false
false

用到的相关的java容器库：
Set<Character> mySet = new HashSet<Character>();
mySet.add('0');
Character s = '0';
mySet.contains(s);

String s = new String("myhello");
s.substring(0,4); ->(0, 4];

面试题21. 调整数组顺序使奇数位于偶数前面
class Solution {
    public int[] exchange(int[] nums) {
        if(nums.length == 0)return nums;
        int left = -1;
        int right = nums.length;
        while(left<right){
            while(++left<nums.length && nums[left]%2 == 1 );
            while(--right>=0 && nums[right]%2 == 0);
            if(left < right){//不然换过去的可能又换了回去
                int tmp = nums[left];
                nums[left] = nums[right];
                nums[right] = tmp;
            }
        }
        return nums;
    }
}
这一题的整个遍历的循环与快排算法中的相同：只是这个遍历有可能会出现终点在边界，也有可能出现终点在中间；而快排不可能出现终点在边界的情况。
注意检查一下快速排序问题



面试题24. 反转链表
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode header = new ListNode(0);
        header.next = null;
        ListNode start = head;
        while(start != null){
            ListNode tmp = start.next;
            start.next = header.next;
            header.next = start;
            start = tmp;
        }
        return header.next;
    }
}





class Solution {
    class node{
        node(int x, int y){
            this.x = x;
            this.y = y;
        }
        int x = 0;
        int y = 0;
    }
    public int[] spiralOrder(int[][] matrix) {
        
        int [] rval = new int[(matrix.length-1)*(matrix[0].length-1)];
        node leftUp = new node(0, 0);
        node rightUp = new node(matrix[0].length-1, 0);
        node leftDown = new node(0, matrix.length-1);
        node rightDown = new node(matrix[0].length-1, matrix.length-1);

        int count = 0;
        for(;;){
            for(int i = leftUp.x; i<rightUp.x; i++){
                rval[count++] = matrix[leftUp.y][i];
                if(count == (matrix.length-1)*(matrix[0].length-1)-1)
                    break;
            }
            for(int i = rightUp.y; i<rightDown.y; i++){
                rval[count++] = matrix[y][leftUp.x];
                if(count == (matrix.length-1)*(matrix[0].length-1)-1)
                    break;
            }
            for(int i = rightDown.x; i>leftDown.x; i--){
                rval[count++] = matrix[rightDown.y][i];
                if(count == (matrix.length-1)*(matrix[0].length-1)-1)
                    break;
            }
            for(int i = leftDown.y; i>leftDown.y; i--){
                rval[count++] = matrix[i][leftDown.x];
                if(count == (matrix.length-1)*(matrix[0].length-1)-1)
                    break;
            }
            leftUp.x;
        }

    }
}

面试题29. 顺时针打印矩阵递归着，一圈一圈的打印，但是要注意终止条件
面试题38. 字符串的排列
就是一个数学中的排列组合：可以开个record数组记录已经排过序的字符，并且可以使用set来进行减枝

面试题40. 最小的k个数
插入算法和快速排序算法都写的不熟练
class Solution {
    public void swap(int [] arr, int indexA, int indexB){
        int tmp = arr[indexA];
        arr[indexA] = arr[indexB];
        arr[indexB] = tmp;
    }

    int getProvit(int []arr, int start, int end){
        int begin = start;
        int finish = end - 1;
        int mid = (start + end)/2;
        if(arr[begin] > arr[finish])
            swap(arr, begin, finish);
        if(arr[mid] < arr[begin])
            swap(arr, mid, begin);
        if(arr[mid]>arr[finish])
            swap(arr, mid, finish);
        return arr[mid];
    }
    public void quickSort(int []arr, int start, int end){
        if(end - start > 3){

            int provit = getProvit(arr, start, end);
            swap(arr, (start+end)/2, end-2);
            int i = start;
            int j = end - 2;
            while(true){
                while(arr[++i]<provit);//等于号是为了解决所有数字都相同的问题
                while(arr[--j]>provit);//快速排序算法中的运算逻辑
                if(i<j){
                    swap(arr, i, j);
                }else{
                    break;
                }
            }
            swap(arr, i, end-2);
            quickSort(arr, start, i);
            quickSort(arr, i+1, end);
        }else{
            insertSort(arr, start, end);
        }
    }
     public void insertSort(int [] arr, int start, int end){
        
        
        for(int i = start; i<end - 1; i++){
            int tmpMin = arr[i];
            int tmpMinIndex = i;
            for(int k = i+1; k<end; k++){
                if(tmpMin > arr[k]){
                    tmpMin = arr[k];
                    tmpMinIndex = k;
                    swap(arr, i, k);
                }
            }
        }
        
        return;
    }
    public int[] getLeastNumbers(int[] arr, int k) {
        //采用快排算法
        int [] rval = new int[k];
        quickSort(arr, 0, arr.length);
        for(int i = 0; i<k; i++){
            rval[i] = arr[i];
        }
        return rval;
    }
}

面试题41. 数据流中的中位数（没有通过）自己使用的LinkedList 加上二分查找法，仍然是超时
正确的应该是使用  优先队列 来实现

面试题42. 连续子数组的最大和
开数组使用动态规划

面试题46. 把数字翻译成字符串使用回溯的方法面试题

48. 最长不含重复字符的子字符串
//dp、滑动窗口、双指针这几种算法都有，自己使用了HashSet