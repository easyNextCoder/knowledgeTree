#include <iostream>
#include <bitset>
using namespace std;

int main()
{
	int a = 1;
	int b = 2;
	cout<<"a^b is: "<<(a^b)<<endl;
	cout<<"a|b is: "<<(a|b)<<endl;
	
	cout<<"(b-a == b+a) is: "<<(b-a == b+a)<<endl; //算术运算符         >   逻辑运算符 
	
	cout<<"(a^b == a^b) is: "<<(a^b == a^b)<<endl;//逻辑运算            >   位运算符 
	
	cout<<(b^b || a^b)<<endl; 					  //位运算符            >  关系运算符 
	return 0;
}
