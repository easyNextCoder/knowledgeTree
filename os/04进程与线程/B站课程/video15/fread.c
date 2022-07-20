//#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <stdio.h>
#include <fcntl.h>


int main(char argc, char** argv ){

	FILE* src_fd, *des_fd;
	int read_ret;
	char buf[128] = {0};
	src_fd  = fopen(argv[1],"r");
	if(src_fd == NULL){
		printf("we can not open file.\n");
	}
	des_fd = open(argv[2],"w+");
	if(des_fd == NULL){
		printf("we can not open des_file.\n");
	}
	while(1){
		read_ret = fread(buf, 1, 128, src_fd);
//		if(feof(fp)){
		if(read_ret < 128){
			fwrite(buf, 1, read_ret, des_fd);
			printf("read file end.\n");
			break;
		}
		fwrite(buf, 1, 128, src_fd);
	}
	close(src_fd);
	close(des_fd);
	return 0;
}
