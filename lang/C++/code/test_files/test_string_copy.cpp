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
	initFromCharp2 = p;	//�ַ�����stringת���Ľӿ� 
	name.copy(ap, 15, 0);//string���ַ���ת�Ľӿ� 
	cout<<name[0]<<endl;
	cout<<ap<<endl;
	cout<<initFromCharp<<endl;
	
	//������Ϊ�˲��ԣ�string��char*֮��Ľӿ� 
	return 0;
}
