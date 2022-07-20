#include <iostream>

using namespace std;


class base{
	public:
		virtual get(){}; 
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
		char ac;//涉及到内存对齐！ 
		int bi;
};

void get(base a){
	cout<<"进行了隐式的类型转换。"<<endl;
};
int main(){
	base * b = new base();
	derived * db = new derived();
	get(*db);
	cout<<"在本机上，int型数据占用多少字节："<<sizeof(int)<<endl;
	cout<<"普通不带虚函数的类占的内存大小是："<<sizeof(*b)<<endl;
	cout<<"普通不带虚函数的内有类对象的类占的内存大小是："<<sizeof(*db)<<endl;;
	 
	delete b;//这里重载了delete操作符 
	delete db;//这里并没有继承重载的delete操作符 
	return 0;
}
