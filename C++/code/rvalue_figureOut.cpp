#include <iostream>
using namespace std;

class RightRef{
public:
	RightRef():b(new int()){
		cout<<"constructor"<<endl;
	}
	RightRef(RightRef& input){
		//b = in.b;
		cout<<"RightRef(&)"<<endl;
	}
	RightRef(RightRef&& input){
		//in.b = nullptr;
		cout<<"invoke RightRef(&&)"<<endl;
	}
private:
	int * b;
};
RightRef getRightRef(){
	RightRef a;
	return a;
}
int main(){
	Right a = getRightRef();
	//在这个过程中，本来应该调用移动构造函数，但是由于现当今的编译器都做了优化，直接把堆中的对象给了目的变量 
	return 0;
}
