#include <iostream>
#include <stdio.h>
#include <string>
#include <string.h> 
using namespace std;

int main()
{
	char*p = "yourname";
	string cp = "yourname";
	char as[] = {'y','o','u','r','n','a','m','e'};
	char pa[] = "yourname";
	cout<<strlen(p)<<endl;
	cout<<strlen(cp.c_str())<<endl;
	cout<<"sizeof char as[]:"<<sizeof(as)<<endl;	
	cout<<"sizeof char pa[]:"<<sizeof(pa)<<endl;
	return 0;
}
