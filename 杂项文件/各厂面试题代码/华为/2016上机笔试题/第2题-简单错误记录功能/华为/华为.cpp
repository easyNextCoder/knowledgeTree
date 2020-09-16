//https://www.nowcoder.com/test/question/done?tid=33849565&qid=25368#summary
//https://www.nowcoder.com/questionTerminal/67df1d7889cf4c529576383c2e647c48
#include <iostream>
#include <sstream>
#include <string>
#include <set>
#include <map>
#include <vector>
#include <algorithm>
using namespace std;


int main()
{
	string str;
	vector<pair<string, int>> vec;
	while (getline(cin, str))
	{
		stringstream ss(str);
		string file;
		int lineNo;
		ss >> file >> lineNo;


		int i = 0;
		for (i = file.size() - 1; i >= 0; i--)
		{
			if (file[i] == '\\')
				break;
		}
		//string fileName = file.substr(i + 1, file.size() - 1 - i);
		string fileName = file.substr(i + 1);
		//上述替换：string.substr(i);默认从当前i值取到字符串的末尾
		vec.push_back(make_pair(fileName, lineNo));
	}

	vector<pair<string, int>>mp;
	map<string, int>con;
	for (auto item : vec)
	{
		string sindex = item.first;
		sindex += " ";
		sindex += to_string(item.second);
		
		if (mp.empty())
		{
			con[sindex] = 0;
			mp.push_back(make_pair(sindex, 1));
		}
		else {
			if (con.count(sindex))
			{
				mp[con[sindex]].second++;
			}
			else {
				con[sindex] = mp.size();
				mp.push_back(make_pair(sindex, 1));
			}
		}
	
	}

	sort(mp.begin(), mp.end(), [](pair<string, int> a, pair<string, int> b)->bool {return a.second > b.second; });
	int idx = 0;
	while (idx < 8 && idx < mp.size()) {
		string check = mp[idx].first;
		int t = check.find(' ');
		if (t > 16)
			mp[idx].first.erase(0, t - 16);
		cout << mp[idx].first << ' ' << mp[idx].second << endl;
		idx++;
	}

}


