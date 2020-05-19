#include <iostream>

using namespace std;

int main()
{
	int i = 1l;
	int *  const p = &i;
	p++;
	printf("%d\n", *p);
	
	
	int* j = new int[2];
	j[0] = 1;
	j[1] = 2;
	int const * q = j;
	q++;
	printf("%d\n", *q);
	
	return 0;
}
