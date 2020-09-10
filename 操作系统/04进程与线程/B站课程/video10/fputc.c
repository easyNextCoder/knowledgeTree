#include "stdio.h"

int main(){
	FILE *fp;
	fp = fopen("./a.c","w+");
	if(fp == NULL){
		printf("open file a.c failure\n");
		return -1;
	}
	printf("open file a.c success.\n");
	fputc('a', fp);
	fputc('\n',fp);
	fflush(fp);
	while(1);

	fclose(fp);
	return 0 ;
	
}
