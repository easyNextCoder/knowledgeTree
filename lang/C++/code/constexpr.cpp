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
	//constexpr int b = a;//��ͨ���������ڱ���ʱ�̾��ǳ��� 
	constexpr int c = 4;
	constexpr int d = xsize();
	constexpr int e = ysize(6);
	//constexpr int e = ysize(c+3);//��ͨ����Ϊ���Ǳ���ʹ��ĳ��� 
	
	const int f = a;//����������ʱ��ʼ������ 
	cout<<"a: "<<a/*<<"b: "<<b*/<<"c: "<<c<<"d: "<<d<<endl;
	return 0;
}
