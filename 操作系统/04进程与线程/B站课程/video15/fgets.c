#include <stdio.h>
#include <stdlib.h>

int main(char argc, char** argv ){

	int return_char;
	FILE *fp = fopen(argv[1], "r");
	char readbuf[128] = {0};
	if(fp == NULL){
		printf("we can not open file.\n");
	}
	FILE *fp_des = fopen(argv[2], "w+");
	if(fp_des == NULL){
		printf("we can not open des_file.\n");
	}
	while(1){
		fgets(readbuf, 128, fp);
		if(feof(fp)){
		//if(return_char < 0){
			printf("read file end.\n");
			break;
		}
		fput(readbuf, fp_des);
	}
	fclose(fp);
	fclose(fp_des);
	return 0;
}
