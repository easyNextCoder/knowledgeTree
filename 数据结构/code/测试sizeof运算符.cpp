#include <iostream>

using namespace std;

int main()
{	
	int b = 0;
	int a = sizeof(b++);
	cout<<"b after sizeof(b++):"<<b<<endl;
	
	char s[] = "yourname";
	char (*ps)[9] = &s;
	cout<<"字符数组的长度是："<<sizeof(s)<<"字符数组指针的大小"<<sizeof(ps)<<endl; 
	return 0;
}
