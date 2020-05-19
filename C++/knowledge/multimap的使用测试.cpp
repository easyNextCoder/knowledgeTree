#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
	multimap<int, string> umap;
	
	umap.insert(make_pair(0,"it"));
	umap.insert(make_pair(0,"is"));
	umap.insert(make_pair(1,"hello"));
	umap.insert(make_pair(1,"what is"));
	umap.insert(make_pair(2,"xu"));
	umap.insert(make_pair(2,"yongkang"));
	
	cout<<"output all value that key is 1"<<endl;
	int num = umap.count(1);
	auto it = umap.find(1);
	for(int i = 0; i<num; ++i)
	{
		cout<<(it->second)<<endl;
		it++;
	} 
	
	cout<<"output all value that key is 2"<<endl;
	auto iter1 = umap.lower_bound(2);
	for(iter1; iter1 != umap.upper_bound(2); ++iter1)
	{
		cout<<(iter1->second)<<endl;
	}
	return 0;
}
