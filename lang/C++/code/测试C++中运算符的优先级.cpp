#include <iostream>
#include <bitset>
using namespace std;

int main()
{
	int a = 1;
	int b = 2;
	cout<<"a^b is: "<<(a^b)<<endl;
	cout<<"a|b is: "<<(a|b)<<endl;
	
	cout<<"(b-a == b+a) is: "<<(b-a == b+a)<<endl; //���������         >   �߼������ 
	
	cout<<"(a^b == a^b) is: "<<(a^b == a^b)<<endl;//�߼�����            >   λ����� 
	
	cout<<(b^b || a^b)<<endl; 					  //λ�����            >  ��ϵ����� 
	return 0;
}
