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
			iter = vec.erase(iter);//map�е�Ԫ��ɾ��֮�󲢲�֧�ָ�ֵ��ɾ��֮��+1λ�õĵ�������Ȼ���� 
		else
			++iter; 
	}
	cout<<"��ⰴ��һ������ɾ��Ԫ�أ�"<<endl; 
	for(auto item: vec)
	{
		cout<<item<<endl;
	}
	cout<<"����ȫ��ɾ��Ԫ�أ�"<<endl;
	iter = vec.begin();
	while(iter != vec.end())
	{
		iter = vec.erase(iter);
	}
	
	if(vec.empty())
		cout<<"ȫ��ɾ����ϣ�"<<endl;
	
	return 0;	
} 
