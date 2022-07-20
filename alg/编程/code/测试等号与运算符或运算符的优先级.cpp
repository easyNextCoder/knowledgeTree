#include <iostream>

using namespace std;

int main()
{
	int L = 1;
	int M = 4;
	int N = 2;
	int result = 0;
	result =( (L | N)& M );
	/*
	0111
	0001 L
	0010 M
	0100 N
	*/
	cout<<result<<endl;
	return 0;
}

