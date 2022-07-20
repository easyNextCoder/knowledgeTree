#include "unistd.h"
#include "sys/types.h"
#include "stdio.h"
#include "string.h"
#include "fcntl.h"
#include "fcntl.h"
#include "dirent.h"
int main(){

	DIR *dp;
	int fd;
	struct dirent* dir;
	char server[128] = {0};
	char file[128] = {0};
	int src_fd = 0;
	int des_fd = 0;
	char buffer[128] = {0};	
	printf("please input server Path and Directory name.\n");
	scanf("%s", server);
	while( (dp = opendir(server)) == NULL){
		printf("open server:%s failure\n", server);
		printf("please input server Path and Directory name.\n");
		scanf("%s", server);
	}	
	
	while(1){
		
		dir = readdir(dp);
		if(dir == NULL){
			break;
		}else{
			printf("inode = %ld file name = %s\n", dir->d_ino, dir->d_name);
		}
	}
	
	printf("\n Please input download file name.\n");
	scanf("%s",file);	

	src_fd = open(strcat(strcat(server, "/"),file), O_RDONLY);
	if( src_fd < 0 ){
		printf("open downloading file:%s\n",file);
		return -1;
	}
	printf("open download file:%s\n",file);
	des_fd = open(file,O_CREAT | O_WRONLY, 0777 );
	//printf("des_fd %d\n", des_fd);
	if(des_fd < 0){
		printf("create file:%s failure.\n", file);
		return -2;
	}	
	printf("create file:%s.\n",file);
	int buffer_return;
	int buffer_write_return;
	int min_buffer_size = 128;
	while(1){

			buffer_return  = read(src_fd, buffer, min_buffer_size);
			if(buffer_return == 0){
				break;	
			}else{
				min_buffer_size = (buffer_return > min_buffer_size)?min_buffer_size:buffer_return;
			}
			buffer_write_return = write(des_fd, buffer, min_buffer_size);
	}
	printf("download OK.\n");
	closedir(dir);
	return 0;
}


