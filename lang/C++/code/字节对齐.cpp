#include <iostream>

using namespace std;

class a{
	char b;
	double c;
	int d;
};
//最终字节数要是sizeof(double) 8的倍数 

class b{
	double c;
	char d;
	int e;
};

class c{
	double d;
	int e;
	char f;
}; 
//https://www.cnblogs.com/lazycoding/archive/2011/03/22/sizeof-struct.html


typedef struct _tag_PARAM {
int ia;
char  cb;
char  cc;
int id;
char  ce;
} Pa;

int main()
{
	cout<<sizeof(a)<<endl;
	cout<<sizeof(b)<<endl;
	cout<<sizeof(c)<<endl;
	cout<<sizeof(Pa)<<endl;
	return 0;
}
