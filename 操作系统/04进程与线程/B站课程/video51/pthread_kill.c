#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>
void* thread_fun(void *arg){
	
	sleep(1);
	printf("I am new thread.\n");
	return (void *)0;
}

int main(){

	pthread_t tid;
	int err;
	int s;
	void *rval;
	
	err = pthread_create(&tid, NULL, thread_fun, NULL);
	if(err != 0){
		printf("we can not create new pthread.\n");
	}

	s = pthread_kill(tid, SIGQUIT);
	if(s != 0){
		printf("thread tid is not found, return value is:%d\n",s);
	}
	
	pthread_join(tid, &rval);
	printf("I am man thread.\n");
	return 0;
}
