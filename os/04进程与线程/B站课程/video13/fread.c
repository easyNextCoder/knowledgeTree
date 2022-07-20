#include <stdio.h>
#include <stdlib.h>

int main(char argc, char** argv ){

	char buf[] = "hello linux.\n";
	char read_buf[128] = {0};
	FILE *fp = fopen("a.c", "w+");
	if(fp == NULL){
		printf("we can not open file.\n");
	}
	printf("open file a.c success.\n");
	fwrite(buf, sizeof(char),sizeof(buf),fp);
	rewind(fp);	
	fread(read_buf, sizeof(char), sizeof(read_buf), fp);
	printf("read_buf:%s\n", read_buf);
	fclose(fp);
	return 0;
}
