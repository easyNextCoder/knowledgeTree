#include <stdio.h>

struct Test{};
struct Test1{
	enum month{Janurary, February};
	enum day{Monday, Tuesday};
};
enum year{y2019, y2020 = 2020};
int main()
{
	struct Test c;
	struct Test1 c1;
	printf("��struct�����Ĵ�С:%d\n", sizeof(c));
	printf("struct���ں�enum�����Ĵ�С:%d", sizeof(c1));
	
	enum year nowy = y2020;
	printf("now year value is:%d\n", nowy);
	enum year lasty = y2019;
	printf("last year value is:%d\n", lasty);
	
	cout<<sizeof(string)<<endl;
	
	return 0;
} 
