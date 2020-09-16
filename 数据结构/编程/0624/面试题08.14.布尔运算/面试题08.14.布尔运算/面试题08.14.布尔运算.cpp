// 面试题08.14.布尔运算.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <string>
#include <set>

using namespace std;

class Solution {
public:
	set<string> out;
	string sr;
	void dfs(int index, int leftbs, int rightbs, int maxbs, string expr)
	{
		//大于最多有效括号数跳出
		if (maxbs > sr.size() - 1)
		{   //加的有用的括号的最大数量是maxbs
			return;
		}

		//终止条件是括号都用完了
		if (index == sr.size())
		{
			if (rightbs == 0)
			{   //此时要么一个括号没有要么是已经加上了有效的括号个数
				out.insert(expr);
			}
			else if (rightbs == 1)
			{
				expr.push_back(')');
				out.insert(expr);
			}
			return;
		}

		//遇到了运算符直接加入表达式并进入下一个
		
		if (sr[index] == '^' || sr[index] == '|' || sr[index] == '^')
		{
			expr.push_back(sr[index]);
			dfs(index + 1, leftbs, rightbs, maxbs, expr);
		}
		else {

			

			//加左括号
			//后表达式不加括号
			expr.push_back(sr[index]);
			dfs(index + 1, leftbs, rightbs, maxbs, expr);

			//先表达式加括号
			expr.pop_back();
			expr.push_back('(');
			expr.push_back(sr[index]);
			dfs(index + 1, leftbs, rightbs + 1, maxbs + 1, expr);

			//加右括号
			if (rightbs > 0)
			{
				//不加右括号
				dfs(index + 1, leftbs, rightbs, maxbs + 1, expr);
				//加右括号
				expr.push_back(')');
				dfs(index + 1, leftbs, rightbs - 1, maxbs + 1, expr);
			}
		}	

	}

	int countEval(string s, int result) {
		int maxbs = s.size() - 1;
		sr = s;
		string expr = "";
		dfs(0, 0, 0, 0, expr);
		for (auto item : out)
			cout << item << endl;
		return 0;
	}
};



int main()
{
	Solution so;
	string s = "1^0|0|1";
	int result = 0;
	so.countEval(s, result);
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




