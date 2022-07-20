#include <stdio.h>

int main(){
	char s[10];

	while( scanf("%s",s) != EOF)
		printf("the input s is:%s\n", s);
	printf("we are exiting.\n");
	return 0;
}
