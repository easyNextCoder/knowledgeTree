#include <iostream>

using namespace std;

int main()
{
	
	int a = -1;
	int b = 1; 
	cout<<"�з������ε���λ������"<<(a<<1)<<endl;
	cout<<"�з������ε���λ������"<<(b>>1)<<endl;
	int aa = INT_MIN+2;
	cout<<"�з������ε���λ������"<<(aa<<1)<<endl;//
	//�з������α���������λ�������ǵ�ʱ���ѷ���λ���� 
	//������λʱ��Ҫ���в�λ������������ʱ�򣬲�0�����Ǹ�����ʱ��1 
	
	unsigned int c = 1;
	unsigned int d = -1;
	cout<<"�޷������ε���λ������"<<(c<<1)<<endl;
	cout<<"�޷������ε���λ������"<<(d<<0)<<endl;
	cout<<"�޷������ε���λ������"<<(d<<1)<<endl;//�ǲ�0�� 
	
	int arr[3][3];
	int (*p)[3];
	p = arr;
	cout<<p[1][0]<<endl;
	
	char s[] = "abcde";
	s += 2;
	printf("%c\n", s[0]);
	
	return 0;
}
