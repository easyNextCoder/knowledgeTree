// vivo2020届春季校园招聘在线编程-1.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//nowcoder.com/questionTerminal/c552248efdbd41a18d35b7a2329f7ad8

#include <iostream>
#include <algorithm>
#include <unordered_set>
#include <vector>
#include <set>

using namespace std;

vector<int> result(10, 0);
class Solution {
public:
	/**
	 * 实现方案
	 * @param m int整型 最少m个键
	 * @param n int整型 最多n个键
	 * @return int整型
	 */
	unordered_set<int> p = { 1,3,7,9 };
	unordered_set<int> v = { 2,4,6,8 };
	void helper2(vector<char>& arr, int used, int levelCount)
	{
		for (int i = 1; i <= 9; i++)
		{
			if (arr[i])continue;
			if ( (p.count(used) && p.count(i)) ||//从顶点出发到顶点可能会发生覆盖
				 ((used == 2 && i == 8) ||(used == 4 && i == 6) ||(used == 6 && i == 4) ||(used == 8 && i == 2))
				 //从边的中心出发可能出现覆盖的情况
				)
			{
				//代表中间会被覆盖的点已经访问过了
				if (arr[(used + i) / 2] == 2)
				{
					arr[used] = 2;
					arr[i] = 1;
					helper2(arr, i, levelCount + 1);
					arr[i] = 0;
					arr[used] = 1;
					result[levelCount]++;
				}
			}
			else {
				arr[used] = 2;
				arr[i] = 1;
				helper2(arr, i, levelCount + 1);
				arr[used] = 1;
				arr[i] = 0;
				result[levelCount]++;
			}
		}

	}
	int solution(int m, int n) {
		// write code here
		int levelCount = 0;
		vector<char> arr(10, 0);// = { 0 };

		for (int i = 1; i <= 9; i++)
		{
			arr[i] = 1;
			helper2(arr, i, levelCount + 1);
			arr[i] = 0;
			result[levelCount]++;
		}

		int rval = 0;
		n = n > 9 ? 9 : n;
		for (int i = m - 1; i < n; i++) {
			rval += result[i];
		}
		return rval;
	}
};


int main()
{
	Solution solution;
	cout << solution.solution(5, 5);
	//std::cout << "Hello World!\n";
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
