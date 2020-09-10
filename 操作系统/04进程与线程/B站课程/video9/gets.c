#include "stdio.h"
int main(){
	char buf[128] = {0};
	int len;
	gets(buf);//have \n
	len = strlen(buf);
	printf("len = %d\n", len);
	fputs(buf, stdout);
	return 0;
}
