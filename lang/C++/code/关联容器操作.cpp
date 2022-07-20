#include <iostream>
#include <set>
#include <map>

using namespace std;

int main()
{
	
	cout<<"test set."<<endl;
	set<int> mst;
	mst.insert(1);
	mst.insert(2);
	mst.insert(3);
	mst.erase(3);
	cout<<"测试插入重复元素是否成功："<<mst.insert(1).second<<endl;
	set<int>::iterator mst_iter = mst.begin();
	
	while(mst_iter != mst.end())
	{
		cout<<*mst_iter++<<endl;
	}
	
	return 0;
}
