#include <iostream>


using namespace std;

template<class T>
struct miterator_traits{
	typedef T value_type;
};

//这个就是对原类型进行偏特化
//如果没有这个偏特化输出值是地址形式，证明偏特化起了作用 
template <class T>
struct miterator_traits<const T*>
{
	typedef T value_type;
};

int main()
{
	
	miterator_traits<const int*>::value_type a;
	a = 10;
	cout<<a<<endl;
	
	return 0;
}
