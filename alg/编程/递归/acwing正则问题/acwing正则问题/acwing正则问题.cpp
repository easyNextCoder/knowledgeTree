#include <iostream>
#include <string>
#include <algorithm>

using namespace std;
string str;
int u = 0;
int n = 0;

/*
从这个回溯中自己得到的一些收获：
1. 首先对于函数来讲：如何界定一次回溯要完成的功能
	对于此题来讲一次回溯就是处理完一对()号中的内容
2. 首先是回溯的过程中元素的推进方式

3. 什么时间递归进入下一级
	对于本题来讲是碰到(时候
4. 什么时间退出此级递归函数
	对于本题来讲是碰到)时候
5. 退出此级函数之后，接着上一级函数的断点来运行
*/

int dfs()
{
	int thisLevelxNum = 0;
	int storeNum = 0;
	if (u == n)return 0;
	while (u < n)
	{
		if (str[u] == '(')
		{
			u++;
			thisLevelxNum = thisLevelxNum + dfs();
		}
		else if (str[u] == 'x')
		{
			thisLevelxNum++;
			u++;
		}
		else if (str[u] == '|') {
			u++;
			storeNum = thisLevelxNum;
			thisLevelxNum = 0;//先保存左边的内容，再开始记录|右边的内容
		}
		else if (str[u] == ')') {
			u++;
			break;
		}
	}
	return max(storeNum, thisLevelxNum);
}

int main()
{

	string instr;
	cin >> instr;
	str = "(";
	str += instr;
	str += ")";
	int start = 0;
	n = str.size();
	cout << dfs() << endl;

	return 0;
}