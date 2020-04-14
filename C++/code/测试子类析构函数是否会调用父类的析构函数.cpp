#include <iostream>

using namespace std;

class A{
	public:
		int ai = 0;
		virtual ~A() = 0;
};
A::~A(){cout<<"invoking ~A()"<<endl;}
//类中生命了纯虚函数之后就是，抽象类了，就算有函数实体，也不能实例化；但是如果被用来继承就必须有函数体可供调用。 

class B:public A{
	public:
		int bi = 0;
		~B(){
			cout<<"invoking ~B()"<<endl;
		}
};




int main(){
	
	
	//当A中的析构函数不是虚函数的时候释放a的时候只会调用~A(); 
	A* b = 	new B();
	//b->bi;
	delete b;
	
	return 0;
}
