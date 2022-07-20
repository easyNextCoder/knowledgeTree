//����Ż��汾��swap()��Ҳ��stl�汾���е�ʵ����ʽ 
//���е�STL�������ṩ��public swap��Ա������
//std::swap�ػ��汾�����Ե���ǰ�ߣ� 

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
			using std::swap;//��ʹ����һ�п��ܻᷢ��ѭ�����û����Ҳ�����Ӧ�ĺ��� 
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


/*����һ�ִ����������ǲ��޷����ɣ���ģ���swap���� 
namespace std{//��std�ռ� 
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


