#include <iostream>

using namespace std;

struct BC{
	unsigned char a:1;
	unsigned char b:1;
	unsigned char c:1;
	unsigned char d:1;
	unsigned char e:1;
	unsigned char f:1;
	unsigned char g:1;
	unsigned char h:1;
	
	
	
	
}; 

int main()
{
	BC a;
	char c = 0;
	BC* pc = new(&c) BC();
	 
	a.a = 0;
	a.b = 0;
	a.c = 0;
	a.d = 0;
	a.e = 0;
	a.f = 0;
	a.g = 1;
	a.h = 0;
	
	pc->a = 1;
	cout<<sizeof(BC)<<endl; 
	cout<<c<<endl;
	return 0;
}
