#include <iostream>

using namespace std;

class FontHandle{
	
};


class Font{
public:
	operator FontHandle(){//这里就很奇怪，到底是重载了什么东西？ 
		return fh;
	}	
	void operator()()
	{
		cout<<"using () override."<<endl;
	}
	/*
		operator一共两种用法一个是第一种表示隐式类型转换，第二种表示重载操作符 
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
