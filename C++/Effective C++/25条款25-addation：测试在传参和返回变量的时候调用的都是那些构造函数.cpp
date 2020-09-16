#include <iostream>
#include <vector>
#include <string>

using namespace std;

//任何函数当其参数对象是：引用类型的时候，不需要调用拷贝构造函数
//当函数的返回值类型是：引用类型的时候也不用调用拷贝构造函数，
					  //但是当其返回的类型是非引用类型的时候就需要调用拷贝构造函数 

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
