#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

void* thread_fun(void *arg){
	//如果传入的参数是1就使用return方式退出
	if(strcmp("1",(char*)arg) == 0){
		printf("new thread return!\n");
		return (void*)1;
	}
	//如果传入参数是2就使用pthread_exit方式退出
	if(strcmp("2",(char*)arg) == 0){
		printf("new thread pthread_exit!\n");
		pthread_exit((void*)2);
	}
	//如果传入参数3，就采用exit方式退出
	if(strcmp("3",(char*)arg) == 0){
		printf("new thread exit!\n");
		exit((void*)3);
	}
	
}

int main(int argc, char* argv[]){
	
	int err;
	pthread_t tid;
		
	err = pthread_create(&tid, NULL, thread_fun, (void*)argv[1]);
	if(err != 0){
		printf("create new thread failed.\n");
		return 0;
	}
	sleep(1);
	printf("main thread\n");
	return 0;
}
