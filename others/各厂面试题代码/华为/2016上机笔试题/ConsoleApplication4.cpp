#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <unordered_map>
#include <algorithm>
//太费劲了，耗费了太多的时间
using namespace std;

bool cmp(pair<string, int> a, pair<string, int> b) { return a.second > b.second; }

int main()
{
	string  s;
	int line;
	map<string, int> con;
	vector<pair<string, int>> tosort;
	while (cin >> s >> line)
	{
		int last = s.size() - 1;
		string filename;
		while (last > 0 && s[last] != '\\')
		{
			filename.push_back(s[last]);
			last--;
		}
		reverse(filename.begin(), filename.end());
		filename.push_back(' ');
		filename += to_string(line);
		if (con.count(filename))
		{
			tosort[con[filename]].second++;
		}
		else {
			con[filename] = tosort.size();
			tosort.push_back({ filename, 1 });
		}

	}


	sort(tosort.begin(), tosort.end(), cmp);
	for (int i = 0; i < min(8, (int)tosort.size()); ++i)
	{
		string tmp = tosort[i].first;
		int npos = tmp.find(' ');
		string name, line;
		line = tmp.substr(npos + 1);

		if (npos > 16)
		{
			tmp = tmp.substr(npos - 16, 16);
		}
		else {
			tmp = tmp.substr(0, npos);
		}
		cout << tmp << " " << line << " " << tosort[i].second << endl;
	}

	return 0;
}