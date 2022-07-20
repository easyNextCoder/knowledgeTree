#include <stdio.h>
#include <stdlib.h>

int main(char argc, char** argv ){

	char buf[] = "hello linux.\n";
	FILE *fp = fopen("a.c", "w+");
	if(fp == NULL){
		printf("we can not open file.\n");
	}
	printf("open file a.c success.\n");
	fwrite(buf, sizeof(char),sizeof(buf),fp);
	while(1);
	//now we can not see text in a.c, so it is not a line cache
	fclose(fp);
	return 0;
}
