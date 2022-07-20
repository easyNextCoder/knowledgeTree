#include <iostream>

using namespace std;


class base{
	public:
		void operator delete(void*s){
			cout<<"重载了delete运算符。"<<endl;
			
			return ;
		}
		
};


class derived{
	public:
		operator base(){
			return obj;
		}
	private:
		base obj;
};

void get(base a){
	cout<<"进行了隐式的类型转换。"<<endl;
};
int main(){
	base * b = new base();
	derived dobj;
	get(dobj);
	
	delete b;
	return 0;
}
