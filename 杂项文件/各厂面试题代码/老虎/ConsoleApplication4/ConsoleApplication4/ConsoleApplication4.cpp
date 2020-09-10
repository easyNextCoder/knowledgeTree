// ConsoleApplication4.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
class Solution {
public:
	/**
	 * 合适的股票每手价格组合
	 * @param prices int整型一维数组 股票数量
	 * @param pricesLen int prices数组长度
	 * @param m int整型 资产值
	 * @return int整型vector<vector<>>
	 */
	void get(int* prices, int priceLen, int m, int index, int tmp,
		vector<int>& tmpv, vector<vector<int>>& res, vector<int>& visited)
	{
		if (index == priceLen)
		{
			if (tmp == m)
			{
				res.push_back(tmpv);
			}
			return;
		}

		get(prices, priceLen, m, index + 1, tmp, tmpv, res, visited);
		if (index > 0 && prices[index] == prices[index - 1] && visited[index - 1] == 0)return;
		tmpv.push_back(prices[index]);
		visited[index] = 1;
		get(prices, priceLen, m, index + 1, tmp + prices[index], tmpv, res, visited);
		tmpv.pop_back();
		visited[index] = 0;
		return;
	}

	vector<vector<int> > combinationSum(int* prices, int pricesLen, int m) {
		// write code here
		if (pricesLen > 15)return {};
		sort(prices, prices + pricesLen);
		vector<vector<int>> res;
		vector<int> tmpv;
		vector<int> visited(pricesLen, 0);
		get(prices, pricesLen, m, 0, 0, tmpv, res, visited);
		vector<vector<int>> rres;
		for (auto& item : res)
		{
			vector<int> tmp;
			for (int i = 0; i < item.size(); ++i)
			{
				if (tmp.empty())
					tmp.push_back(item[i]);
				else if (tmp.back() == item[i])
				{
					continue;
				}
				else {
					tmp.push_back(item[i]);
				}
			}
			rres.push_back(tmp);
		}
		return rres;
	}
};



int main()
{
	Solution so;
	const int N = 10;
	int arr[N] = { 3,7,8,9,10,11,12,13,14,20};

	vector<vector<int>> res = so.combinationSum(arr, N, 20);
	for (auto item : res)
	{
		for (auto line : item)
			cout << line << " ";
		cout << endl;
	}
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



/**
	 * 排队
	 * @param head ListNode类 头结点
	 * @return ListNode类
	 */
	 /**
 * struct ListNode {
 *	int val;
 *	struct ListNode *next;
 * };
 *//*



class Solution {
public:
	
	ListNode* lineUp(ListNode* head) {
		// write code here
		ListNode* headc = head;
		ListNode* headc1 = head->next;
		ListNode* headc2 = head->next;

		while (headc1 && headc1->next)
		{
			headc->next = headc1->next;
			headc1->next = headc1->next->next;
			headc1 = headc1->next;
			headc = headc->next;
		}
		headc->next = headc2;
		return head;
	}
};
*/