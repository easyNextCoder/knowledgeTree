#include <iostream>

using namespace std;

constexpr int xsize()
{
	return 5;
}

constexpr int ysize(int a){
	return a;
}
int main()
{
	int a = 3;
	//constexpr int b = a;//不通过，必须在编译时刻就是常量 
	constexpr int c = 4;
	constexpr int d = xsize();
	constexpr int e = ysize(6);
	//constexpr int e = ysize(c+3);//不通过因为不是编译使其的常量 
	
	const int f = a;//可以在运行时初始化常量 
	cout<<"a: "<<a/*<<"b: "<<b*/<<"c: "<<c<<"d: "<<d<<endl;
	return 0;
}
