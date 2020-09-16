#include "stdio.h"
#include "sys/types.h"
#include "dirent.h"

int main(char argc, char ** argv){

	int ret = 0;
	DIR *dp;
	struct dirent* dir;
	struct dirent* roc;
	//ret = mkdir("./mydir", 0777);
	if(ret < 0)
		printf("create failure.\n");
	printf("create success.\n");
	dp = opendir(argv[1]);
	if(dp ==  NULL){
		printf("open mydir failure.\n");
		return -2;
	}
	while( (dir = readdir(dp)) != NULL){
		roc = telldir(dp);
		printf("roc:%d\n", roc);
		if(dir != NULL){
			printf("%ld %s\n",dir->d_ino,dir->d_name);
		}
		
	}
	printf("open mydir success.\n");
	closedir(dp);
	return 0;
}
