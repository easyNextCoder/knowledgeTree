#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;
vector<pair<int, int>> result;
//返回的是匹配的字符串的长度
void match(int start, int len, string pattern, string str)
{
	if (pattern.empty())
	{
		if (result.empty())
			result.push_back({ start, len });
		else if (!(result.back().first == start && result.back().second == len)) {
			result.push_back({ start, len });
		}
		return;
	}
	if (str.empty())
	{
		return;
	}

	if (pattern[0] == str[0])
	{
		match(start, len + 1, pattern.substr(1), str.substr(1));
	}
	else if (pattern[0] == '*')
	{
		//*匹配当前字母
		match(start, len + 1, pattern.substr(1), str.substr(1));
		//*匹配两个字母
		match(start, len + 1, pattern, str.substr(1));


		//*匹配空
		match(start, len, pattern.substr(1), str);
	}
	else {
		return;
	}
}


int main()
{
	string pattern, str;
	cin >> pattern >> str;
	for (size_t i = 0; i < str.size(); i++)
	{
		match(i, 0, pattern, str.substr(i));
	}
	for (auto item : result)
	{
		cout << item.first << " " << item.second << endl;
	}
	if (result.empty())
		cout << "-1 0" << endl;


	return 0;
}




//算是14.20开始做的吧