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
	multimap<string, int>con = {
		{"xuyongkang", 1},
		{"xuyongqing", 3},

		{"yongqing", 9},
		{"uyongqing", 5},
		{"xuongqing", 3},
		{"xuyongkang", 2}
	};
	multimap<string, int>::iterator iter = con.begin();
	for(;iter!=con.end();iter++)
	{	//这里容器中的元素可以按照顺序输出
		cout<<iter->first<<iter->second<<endl;
	}



	cout << "一个键值对应多个元素的使用方法." << endl;
	multimap<string,int>::iterator  iter_first = con.lower_bound("xuyongkang");
	multimap<string, int>::iterator iter_first_copy = iter_first;
	multimap<string, int>::iterator  iter_last = con.upper_bound("xuyongkang");
	while (iter_first != iter_last)
	{
		cout << iter_first->first << iter_first->second << endl;
		iter_first++;
	}



	cout << "ready to erase key xuyongkang." << endl;
	while (iter_first_copy != iter_last)
	{
		con.erase(iter_first_copy++);
	}
	cout << "after erase  key xuyongkang." << endl;
	for (auto item : con)
	{
		cout << item.first << item.second << endl;
	}
	cout << endl;



	cout << "sort函数只支持random access 接口的迭代器" << endl;
	//sort(con.begin(), con.end(), [](auto a, auto b) {return a.second < b.second; });

	
	return 0;
}
