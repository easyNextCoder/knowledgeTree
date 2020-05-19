#include <iostream>

using namespace std;

int x = 10;
void print(){cout<<"out:we are printing."<<endl;} 
void someFunc(){
	
	double x;
	std::cin>>x;
	cout<<x<<endl;
	print();
	
}


int main(){
	someFunc();
	return 0;
}
