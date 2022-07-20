#include <iostream>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

using namespace std;

char stg[10], str1[10];
int main(){
	*str1 = 'a';
	
	str1[9] = '\0';
	strcpy(stg, str1);	
	return 0;
}
