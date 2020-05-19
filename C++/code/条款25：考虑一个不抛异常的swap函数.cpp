//这个优化版本的swap()，也是stl版本库中的实现形式 
//所有的STL容器都提供有public swap成员函数和
//std::swap特化版本（用以调用前者） 

#include <iostream>
#include <string>
#include <vector>

using namespace std;
namespace WidgetStuff{
	
	template <typename T>
	class WidgetImpl{
	
	public:
		WidgetImpl operator=(const WidgetImpl& a)
		{
			cout<<"invoking WidgetImpl operator="<<endl;	
		}	
		
	private:
		int a, b, c;
		//vector<T> v;
		vector<T>v;
	
	};
	
	template <typename T>
	class Widget{
	
	public:
		Widget() = default;
		
		Widget(const Widget& rhs){
			cout<<"invoking Widget(const Widget& rhs)"<<endl;
		};
		Widget& operator=(const Widget& rhs)
		{
			cout<<"invoking Widget operator="<<endl;
			*pImpl = *(rhs.pImpl);
		}
		
		void swap(Widget& other)
		{
			using std::swap;//不使用这一行可能会发生循环调用或者找不到相应的函数 
			swap(pImpl, other.pImpl);
			cout<<"invoking self swap()."<<endl;
		}
		
	private:
		WidgetImpl<T>* pImpl;
			
	};
	
	template <typename T>
	void swap(Widget<T>& a, Widget<T>& b)
	{
		cout<<"invoking WidgetStuff:void swap(Widget<T>a, Widget<T>b)"<<endl;
		a.swap(b);
		
	}
	

};


/*这是一种处理方法，但是并无法做成，带模板的swap函数 
namespace std{//打开std空间 
	template<>
	void swap<Widget>(Widget& a, Widget& b)
	{
		a.swap(b);
	}
}
*/

using namespace WidgetStuff;





int main()
{
	
	Widget<double> a;
	Widget<double> b;
	swap(a,b);
	
	return 0;
}


