#include <iostream>
#include <stdio.h>

using namespace std;

int main()
{
	char a[20];
	char* p = "yourname";
	char* copy = NULL;
	a[0] = 0;
	//p[0] = '0';
	cout<<p<<endl;
	p = copy;
	cout<<a<<endl;
	//cout<<p<<endl;//������0��ַ��ʱ����ó������ 
	
	cout<<"���������sizeof: "<<sizeof(a)<<endl;
	cout<<"ָ�������sizeof: "<<sizeof(p)<<endl; 
	return 0;
}
