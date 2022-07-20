#include <iostream>

using namespace std;


int z[10];
int main()
{
	int a[10];
	static int b[10];
	
	cout<<"全局变量中的数组的初始化情况："<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<z[i]<<" ";
	}
	cout<<endl;
	
	cout<<"main函数中普通数组的初始化情况："<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<a[i]<<" ";
	}
	cout<<endl;
	
	cout<<"静态数组的初始化情况："<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<b[i]<<" ";
	}
	cout<<endl;
	
	return 0;
}
