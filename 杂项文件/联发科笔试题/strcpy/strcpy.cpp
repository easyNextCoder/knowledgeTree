#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void alloc(char ** a)
{
	*a = (char*)malloc(sizeof(char)*100);
	
}

int main()
{
	char* a;
	alloc(&a);
	//char b[100];
	strcpy(a, "hello world");
	printf("%s", a);
	return 0;	
} 
