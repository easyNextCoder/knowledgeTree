#include <iostream>
#include <string>
#include <vector>

using namespace std;

template<typename T>
class WidgetImpl{

public:
		
private:
	int a, b, c;
	vector<double> v;

};


template<typename T>
class Widget{

public:
	Widget() = default;
	Widget(const Widget& rhs);
	Widget& operator=(const Widget& rhs)
	{
		*pImpl = *(rhs.pImpl);
	}
	
	void swap(Widget& other)
	{
		using std::swap;
		swap(pImpl, other.pImpl);
		cout<<"invoking self swap()."<<endl;
	}
	
private:
	WidgetImpl<T>* pImpl;
		
};

/* 
namespace std{
	//这个是函数的偏特化版本，注意函数无法实现偏特化版本 
	//error:function template partial specializa 
	//这样做是行不通的 
	template<typename T>
	void swap< Widget<T> >(Widget<T>& a, Widget<T>& b)
	{
		a.swap(b);
	}
}
*/

//偏特化版本不能实现想法，所以尝试使用函数重载
namespace std{
	template<typename T>
	swap(Widget<T>a, Widget<T>b)
	{
		a.swap(b);
	}
} 

int main()
{
	
	return 0;	
} 
