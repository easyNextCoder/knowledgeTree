#include <iostream>

using namespace std;

class FontHandle{
	
};


class Font{
public:
	operator FontHandle(){//����ͺ���֣�������������ʲô������ 
		return fh;
	}	
	void operator()()
	{
		cout<<"using () override."<<endl;
	}
	/*
		operatorһ�������÷�һ���ǵ�һ�ֱ�ʾ��ʽ����ת�����ڶ��ֱ�ʾ���ز����� 
	*/
private:
	FontHandle fh;	
}; 

void do_func(FontHandle fha)
{
	cout<<"using do function()."<<endl;
}

int main()
{
	Font f;
	f();
	do_func(f);
	FontHandle f1 ;
	f1 = FontHandle(f);
	return 0;
}
