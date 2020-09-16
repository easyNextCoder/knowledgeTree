//#include <unistd.h>
//#include <sys/types.h>
#include <stdio.h>
//#include <fcntl.h>

int main(int argc, char *argv[]){

	FILE *fp;
	char buf[] = "hello linux\n";
	char read_buf[128] = {0};
	fp = fopen("./a.c", "w+");
	if(fp < 0){
		printf("open file a,c failure.\n");
		return -1;
	}			
	printf("open file a.c success\n");
	fputs(buf, fp);
	fseek(fp, 0, SEEK_SET);
	fgets(read_buf,128, fp);
	printf("the fgets function is:%s\n", read_buf);
	fclose(fp);
    return 0;
}
