#include <iostream>
#include <vector>
#include <string>

using namespace std;

//�κκ���������������ǣ��������͵�ʱ�򣬲���Ҫ���ÿ������캯��
//�������ķ���ֵ�����ǣ��������͵�ʱ��Ҳ���õ��ÿ������캯����
					  //���ǵ��䷵�ص������Ƿ��������͵�ʱ�����Ҫ���ÿ������캯�� 

class Base{

public:
	Base(){
		cout<<"invoking Base()"<<endl;
	}	
	Base(const Base&a)
	{
		cout<<"invoking Base(const Base&a)"<<endl;
	}
	
	Base& test( Base& a)
	{
		cout<<"invoking Base test(const Base& e)."<<endl;
		return *this;
	}

}; 

int main()
{
	Base a;
	Base b;
	cout<<&b<<endl;
	b = b.test(a);
	cout<<&b<<endl;
	
	return 0;
}
