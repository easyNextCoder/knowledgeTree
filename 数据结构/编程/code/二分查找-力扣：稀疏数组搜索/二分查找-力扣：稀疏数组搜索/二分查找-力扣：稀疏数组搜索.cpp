// 二分查找-力扣：稀疏数组搜索.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <string>
#include <vector>
#include <list>

using namespace std;

class Solution {
public:
	int findString(vector<string>& words, string s) {
		if (words.size() == 1)
		{
			if (words[0] == s)return 0;
			else return -1;
		}
		int first = 0;
		int last = words.size() - 1;
		int mid = first + (last - first) / 2;
		while (first <= last)
		{
			mid = first + (last - first) / 2;
			int midc = mid;
			while (words[mid] == "" && mid >= first)mid--;
			if (!(mid >= first))
			{
				first = midc + 1;
			}
			else {
				if (words[mid] > s)
				{
					last = mid - 1;
				}
				else if (words[mid] < s) {
					first = mid + 1;
				}
				else {
					return mid;
				}
			}
		}
		return -1;
	}
};

int main()
{
	Solution so;
	vector<string> ss = { "at", "", "", "", "ball", "", "", "car", "", "","dad", "", "" };//{ "at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""};
	cout << so.findString(ss, "ball") << endl;
	list<int> mylist;
	
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
