#include <iostream>

using namespace std;


int z[10];
int main()
{
	int a[10];
	static int b[10];
	
	cout<<"ȫ�ֱ����е�����ĳ�ʼ�������"<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<z[i]<<" ";
	}
	cout<<endl;
	
	cout<<"main��������ͨ����ĳ�ʼ�������"<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<a[i]<<" ";
	}
	cout<<endl;
	
	cout<<"��̬����ĳ�ʼ�������"<<endl;
	for(int i = 0; i<10; i++)
	{
		cout<<b[i]<<" ";
	}
	cout<<endl;
	
	return 0;
}
