#include <iostream>

using namespace std;

int main()
{
	int i = 0;
	const int j = 1;
	
	const int* p = &i;
	const int* q = &j;
	int *const s = &i;
	int const* t = &j;
	
	int  * const y = &i;
	const int * const x = &j;
	
	cout<<*p<<*q<<*s<<*t<<endl;
	return 0;
}
