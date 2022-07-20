#include <iostream>
#include <vector>

using namespace std;

int main()
{
	vector<int> vec = {1,2,3,4,5,6,7,8,9};
	vector<int>::iterator iter = vec.begin();
	while(iter != vec.end()) 
	{
		if(*iter<5 && *iter>1)
			iter = vec.erase(iter);//map中的元素删除之后并不支持赋值，删除之后+1位置的迭代器仍然可用 
		else
			++iter; 
	}
	cout<<"测测按照一定条件删除元素："<<endl; 
	for(auto item: vec)
	{
		cout<<item<<endl;
	}
	cout<<"测试全部删除元素："<<endl;
	iter = vec.begin();
	while(iter != vec.end())
	{
		iter = vec.erase(iter);
	}
	
	if(vec.empty())
		cout<<"全部删除完毕！"<<endl;
	
	return 0;	
} 
