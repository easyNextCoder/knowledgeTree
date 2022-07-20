#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
	multimap<int, string> umap;
	umap.insert(make_pair(1,"hello"));
	umap.insert(make_pair(1,"what is"));
	auto it = umap.find(1);
	cout<<it->second<<endl; 
	return 0;
}
