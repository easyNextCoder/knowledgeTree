//#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <stdio.h>
#include <fcntl.h>


int main(char argc, char** argv ){

	int src_fd, des_fd;
	int read_ret;
	char buf[128] = {0};
	src_fd  = open(argv[1],O_RDONLY);
	if(src_fd < 0){
		printf("we can not open file.\n");
	}
	des_fd = open(argv[2],O_CREAT | O_WRONLY, 0777 );
	if(des_fd < 0){
		printf("we can not open des_file.\n");
	}
	while(1){
		read_ret = read(src_fd, buf, 128);
//		if(feof(fp)){
		if(read_ret < 128){
			write(des_fd, buf, read_ret);
			printf("read file end.\n");
			break;
		}
		write(des_fd, buf, read_ret);
	}
	close(src_fd);
	close(des_fd);
	return 0;
}
