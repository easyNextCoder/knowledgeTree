//#include <unistd.h>
//#include <sys/types.h>
#include "string.h"
#include <stdio.h>
//#include <fcntl.h>

int main(int argc, char *argv[]){

	FILE *fp;
	char buf[] = "hello linux\n";
	char read_buf[128] = {0};
	char * ret;
	fp = fopen("./a.c", "w+");
	if(fp < 0){
		printf("open file a,c failure.\n");
		return -1;
	}			
	printf("open file a.c success\n");
	fputs(buf, fp);
	fseek(fp, 0, SEEK_SET);
	fgets(read_buf,128, fp);
	printf("readbuf:%s\n", read_buf);
	memset(read_buf, 0, 128);
	ret = fgets(read_buf, 128, fp);
	printf("second readbuf=%s, ret = %p\n", read_buf, ret);
	fclose(fp);
    return 0;
}
