#include <iostream>

using namespace std;

class base{
public:
	base& operator++()
	{
		value++;
		return *this;
	}
	
	base operator++(int)
	{
		value++;
		return *this;//���������Ĭ�ϵĿ������캯�� 
	}
	void set_value(int a)
	{
		value = a;
	}
	int get_value()
	{
		return value;
	}
	
private:
	int value;
	
};

int main()
{
	base o1;
	o1.set_value(1);
	cout<<(++(o1++)).get_value()<<endl;
	cout<<o1.get_value()<<endl;
	return 0;
}
