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
	//cout<<p<<endl;//当访问0地址的时候会让程序崩溃 
	
	cout<<"数组变量的sizeof: "<<sizeof(a)<<endl;
	cout<<"指针变量的sizeof: "<<sizeof(p)<<endl; 
	return 0;
}
