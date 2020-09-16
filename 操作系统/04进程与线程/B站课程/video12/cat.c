#include <stdio.h>
#include <stdlib.h>

int main(char argc, char** argv ){

	int return_char;
	FILE *fp = fopen(argv[1], "r");
	if(fp == NULL){
		printf("we can not open file.\n");
	}
	while(1){
		return_char = fgetc(fp);
//		if(feof(fp)){
		if(return_char < 0){
			printf("read file end.\n");
			break;
		}
		fputc(return_char, stdout);
	}

	return 0;
}
