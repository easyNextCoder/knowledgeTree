#include <iostream>

using namespace std;

class A{
	public:
		int ai = 0;
		virtual ~A() = 0;
};
A::~A(){cout<<"invoking ~A()"<<endl;}
//���������˴��麯��֮����ǣ��������ˣ������к���ʵ�壬Ҳ����ʵ��������������������̳оͱ����к�����ɹ����á� 

class B:public A{
	public:
		int bi = 0;
		~B(){
			cout<<"invoking ~B()"<<endl;
		}
};




int main(){
	
	
	//��A�е��������������麯����ʱ���ͷ�a��ʱ��ֻ�����~A(); 
	A* b = 	new B();
	//b->bi;
	delete b;
	
	return 0;
}
