#include <iostream>


using namespace std;

template<class T>
struct miterator_traits{
	typedef T value_type;
};

//������Ƕ�ԭ���ͽ���ƫ�ػ�
//���û�����ƫ�ػ����ֵ�ǵ�ַ��ʽ��֤��ƫ�ػ��������� 
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
