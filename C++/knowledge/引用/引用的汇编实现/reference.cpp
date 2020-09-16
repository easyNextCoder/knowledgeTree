//#include <iostream>


int main()
{

	int b = -1;
//	int a = -1;
//	a = a + b;
	int& a = b;
	a = a+b;
	return 0;
}
