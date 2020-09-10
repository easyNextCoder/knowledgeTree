#include "stdio.h"
int main(){
	int i = 0;
	char buf[128] = {0};
	sprintf(buf,"i = %d", i);
	printf("buf = %s\n",buf);
	return 0;
}
