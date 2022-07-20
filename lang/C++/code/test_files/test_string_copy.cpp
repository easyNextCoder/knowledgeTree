#include <iostream>
#include <string>

using namespace std; 

int main()
{	
	char *p = "hello,world! I am p";
	char ap[10];
	string name = "hello.";
	cout<<p<<endl;
	string initFromCharp(p);
	string initFromCharp2;
	initFromCharp2 = p;	//字符串向string转换的接口 
	name.copy(ap, 15, 0);//string向字符串转的接口 
	cout<<name[0]<<endl;
	cout<<ap<<endl;
	cout<<initFromCharp<<endl;
	
	//以上是为了测试，string与char*之间的接口 
	return 0;
}
