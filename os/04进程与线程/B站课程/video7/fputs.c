//#include <unistd.h>
//#include <sys/types.h>
#include <stdio.h>
//#include <fcntl.h>

int main(int argc, char *argv[]){

	FILE *fp;
	char buf[] = "hello linux\n";
	
	fp = fopen("./a.c", "w+");
	if(fp < 0){
		printf("open file a,c failure.\n");
		return -1;
	}			
	printf("open file a.c success\n");
	fputs(buf, fp);
	fclose(fp);
    return 0;
}
