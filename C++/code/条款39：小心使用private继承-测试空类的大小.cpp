#include <iostream>

using namespace std;

class empty{
	int func();	
}; 

int empty::func(){
	return 1;
}



class noEmpty:private empty{
	private:
		int a = 0;
		char b = 0;
		char c = 0;
		char d = 0;
		char e = 0;
		char g = 0;
};

int main(){
	cout<<sizeof(empty)<<endl;
	empty one;
	cout<<sizeof(one)<<endl;
	
	noEmpty two;
	cout<<sizeof(two)<<endl;
	return 0;
}
