#include <iostream>

using namespace std;

class base{
public:
	int geta(){
		return a;
	}
	~base(){
		cout<<"调用了base 的析构函数。"<<endl;
	}
private:
	int a = 0;
};

base& get_base(){
	return *(new base);
}

int main(){
	base& ref = get_base();
	cout<<ref.geta()<<endl;
	//造成heap上的内存泄漏 
	base& main_ref = *(new base);
	cout<<main_ref.geta()<<endl;
	//引用仍然造成heap上的内存泄漏 
	base main_obj;
	cout<<main_obj.geta()<<endl;
	//栈上的空间最终没有泄漏	 
	return 0;
} 
