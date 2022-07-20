#include <iostream>

using namespace std;

int main()
{
	
	int a = -1;
	int b = 1; 
	cout<<"有符号整形的移位操作："<<(a<<1)<<endl;
	cout<<"有符号整形的移位操作："<<(b>>1)<<endl;
	int aa = INT_MIN+2;
	cout<<"有符号整形的移位操作："<<(aa<<1)<<endl;//
	//有符号整形变量向左移位发生覆盖的时候会把符号位覆盖 
	//向右移位时，要进行补位，当是正数的时候，补0，当是负数的时候补1 
	
	unsigned int c = 1;
	unsigned int d = -1;
	cout<<"无符号整形的移位操作："<<(c<<1)<<endl;
	cout<<"无符号整形的移位操作："<<(d<<0)<<endl;
	cout<<"无符号整形的移位操作："<<(d<<1)<<endl;//是补0的 
	
	int arr[3][3];
	int (*p)[3];
	p = arr;
	cout<<p[1][0]<<endl;
	
	char s[] = "abcde";
	s += 2;
	printf("%c\n", s[0]);
	
	return 0;
}
