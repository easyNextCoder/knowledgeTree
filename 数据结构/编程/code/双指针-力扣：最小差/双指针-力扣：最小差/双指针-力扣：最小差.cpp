// 双指针-力扣：最小差.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
	int smallestDifference(vector<int>& a, vector<int>& b) {
		sort(a.begin(), a.end());
		sort(b.begin(), b.end());


		int indexA = 0;
		int indexB = b.size() - 1;

		int tmp_min = abs(a[indexA] - b[indexB]);
		while (indexA < a.size() && indexB >= 0)
		{
			//cout << indexA << endl;
			//cout << indexB << endl;
			int tmp_gap = abs(a[indexA] - b[indexB]);
			int moveAHead = tmp_gap;
			int moveBBack = tmp_gap;
			if (indexA + 1 < a.size())
				moveAHead = abs(a[indexA + 1] - b[indexB]);
			if (indexB > 0)
				moveBBack = abs(a[indexA] - b[indexB - 1]);
			//上面三个变量比较确定下一步的方向
			if (min(moveBBack, moveAHead) < tmp_gap)
			{
				if (moveAHead < moveBBack)
				{//两个数组都有余量
					indexA++;
				}
				else if (moveAHead > moveBBack) {
				//两个数组都有余量
					indexB--;
				}
				else if (moveAHead == moveBBack)
				{//只剩其中一个数组有余量
					if (indexA < a.size())
						indexA++;
					else
						indexB--;
				}
				tmp_min = min(moveAHead, moveBBack);
			}
			else{
				return tmp_min;
			}

		}
		return tmp_min;

	}
};

int main()
{
	vector<int> a = { 0 };
	vector<int> b = { 2147483647 };
	Solution solution;
	cout<<solution.smallestDifference(a, b);
    std::cout << "Hello World!\n";
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

Line 11: Char 32 : runtime error : signed integer overflow : -2147483648 - 2147483647 cannot be represented in type 'int' (solution.cpp)
SUMMARY : UndefinedBehaviorSanitizer : undefined - behavior prog_joined.cpp : 21 : 32
