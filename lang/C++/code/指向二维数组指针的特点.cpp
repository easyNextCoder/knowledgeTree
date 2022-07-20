#include <stdio.h>
#include <iostream>

using namespace std;

int main()
{
	int a[3][2] = {{0,1},{2,3},{4,5}};
	int* p = a[1];
	printf("%d\n", p[1]); 
	//cout<<p[0]<<endl;
	
	return 0;
}
