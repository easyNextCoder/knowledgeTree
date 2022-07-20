/*
	data:2020.6.27 更改
	条款25：考虑写出一个不抛异常的swap函数
	在这个条款中我们讨论过default swap, member swap
	non-member swap， std::swap特化版本，以及对swap
	函数的调用

 */

//这个优化版本的swap()，也是stl版本库中的实现形式 
//所有的STL容器都提供有public swap成员函数和
//std::swap特化版本（用以调用前者） 

#include <iostream>
#include <string>
#include <vector>

using namespace std;
namespace WidgetStuff {

	template <typename T>
	class WidgetImpl {

	public:

	private:
		int a, b, c;
		vector<T> v;

	};

	template <typename T>
	class Widget {

	public:
		Widget() = default;
		Widget(const Widget& rhs);
		Widget& operator=(const Widget& rhs)
		{
			*pImpl = *(rhs.pImpl);
		}

		void swap(Widget<T>& other)
		{
			using std::swap;
			swap(pImpl, other.pImpl);
			cout << "invoking Widget swap()." << endl;
		}

	private:
		WidgetImpl<T>* pImpl;

	};

	template <typename T>
	void swap(Widget<T>&a, Widget<T>&b)
	{
		cout << "invoking WidgetStuff:void swap(Widget<T>a, Widget<T>b)" << endl;
		a.swap(b);
	}
};
using namespace WidgetStuff;

/*
//这是一种处理方法，但是并无法做成，带模板的swap函数
//一般情况下这样可以把某个类的swap函数封印到std空间中
//但是函数偏特例化是不允许的
namespace std{//打开std空间
	template<typename T>
	void swap<Widget<T>>(Widget<T>& a, Widget<T>& b)
	{
		a.swap(b);
	}
}
*/



namespace Phone{
	template<typename T>
	class iPhone{
		
	public:
		void swap(iPhone<T>& a)
		{
			cout<<"invoking iPhone swap."<<endl;
			using std::swap;
			swap(this->impl, a.impl);
		}

	private:
		vector<T> cont;
		iPhone<T>* impl;
	};

	template<typename T>
	void swap(iPhone<T>& a, iPhone<T>& b){
		a.swap(b);
	}
};

using namespace Phone;

int main()
{

	Widget<double> a;
	Widget<double> b;
	swap(a, b);
	iPhone<int> iPhone4s;
	iPhone<int> iPhone6s;
	swap(iPhone4s, iPhone6s);
	return 0;
}
//看一看第14章



