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
	//����Ǻ�����ƫ�ػ��汾��ע�⺯���޷�ʵ��ƫ�ػ��汾 
	//error:function template partial specializa 
	//���������в�ͨ�� 
	template<typename T>
	void swap< Widget<T> >(Widget<T>& a, Widget<T>& b)
	{
		a.swap(b);
	}
}
*/

//ƫ�ػ��汾����ʵ���뷨�����Գ���ʹ�ú�������
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
