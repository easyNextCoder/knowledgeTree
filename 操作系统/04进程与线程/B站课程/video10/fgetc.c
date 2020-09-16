#include "stdio.h"

int main(){
	FILE *fp;
	int ret;
	fp = fopen("./a.c","w+");
	if(fp == NULL){
		printf("open file a.c failure\n");
		return -1;
	}
	printf("open file a.c success.\n");
	fputc('a', fp);
	rewind(fp);
	ret = fgetc(fp);
	printf("ret = %c\n",ret);
	ret = fgetc(fp);
	printf("file end is %d\n",ret);//test if EOF is -1
	fclose(fp);
	return 0 ;
	
}
