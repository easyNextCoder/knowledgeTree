#include <iostream>

using namespace std;

int main()
{	
	int b = 0;
	int a = sizeof(b++);
	cout<<"b after sizeof(b++):"<<b<<endl;
	
	char s[] = "yourname";
	char (*ps)[9] = &s;
	cout<<"�ַ�����ĳ����ǣ�"<<sizeof(s)<<"�ַ�����ָ��Ĵ�С"<<sizeof(ps)<<endl; 
	return 0;
}
