// 回溯-布尔运算.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <functional>
#include <map>
#include <vector>

using namespace std;

class Solution {
public:
	typedef function<bool(bool, bool)> func;
	map<char, func>funcMap;

	int count = 0;
	static bool operatorAnd(bool a, bool b)
	{
		return a & b;
	}

	static bool operatorOr(bool a, bool b)
	{
		return a | b;
	}

	static bool operatorXor(bool a, bool b)
	{
		return a ^ b;
	}

	vector<bool> innerCE(string s, int first, int last, int result, int steps)
	{
		cout << "first: " << first << " last: " << last << endl;
		if (first > last)
		{
			cout << "first > last exits." << endl;
			return {};
		}
		if (last - first + 1 >= 3)
		{
			vector<bool> vrval = innerCE(s, first + 2, s.size() - 1, result, steps+1);

			vector<bool> tmpvrval;

			for (auto rv : vrval)
			{
				char calOperator = s[first + 1];
				bool lv = (s[first] == '1');
				bool rval = funcMap[calOperator](lv, rv);
				tmpvrval.push_back(rval);
			}
			vrval.clear();
			vrval = tmpvrval;

			
				vector<bool> vrval1 = innerCE(s, first + 4, s.size() - 1, result, steps+1);
				
				for (auto item : vrval1)
				{
					auto rval1 = funcMap[s[first + 3]](funcMap[s[first + 1]](s[first] == '1', s[first + 2] == '1'), item);
					vrval.push_back(rval1);
				}

				if (vrval1.empty())
				{
					vrval.push_back(funcMap[s[first + 1]](s[first] == '1', s[first + 2] == '1'));
				}
			
			
			if (steps == 0)
			{
				cout << "real size:"<<vrval.size() << endl;
				for(auto item:vrval)
				{ 
					cout << item << " ";
					if (item == result)
						count++;
				}
				cout << endl;
			}
			return vrval;
		}
		else if(last - first + 1 == 3){
			auto tmp = funcMap[s[first + 1]](s[first] == '1', s[first + 2] == '1');
			vector<bool>out;
			out.push_back(tmp);
			return out;
		}
		else {

			if (last - first + 1 == 1)
			{
				cout << "here." << endl;
			}
			return { s[first] == '1' };
		}

	}

	int countEval(string s, int result) {

		//加括号
		//不加括号往前走
		funcMap.insert({ '&', operatorAnd });
		funcMap.insert({ '|', operatorOr });
		funcMap.insert({ '^', operatorXor });

		innerCE(s, 0, s.size() - 1, result, 0);


		return count;
	}
};
int main()
{
	Solution solution;
	cout << solution.countEval("1&0&0|1", 1) << endl;;
	//(1^0|0|1)
	//(1^0
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
