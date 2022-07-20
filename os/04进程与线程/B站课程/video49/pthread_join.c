#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

void* thread_fun1(void *arg){
	
	printf("we are at thread_fun1.\n");	
	return 0;
}
void* thread_fun2(void* arg){
	printf("we are at thread_fun2.\n");
	pthread_detach(pthread_self());
	pthread_exit(2);
}

int main(int argc, char* argv[]){
	
	int err1, err2;
	pthread_t tid1, tid2;
	void *rval1 , *rval2;	
	int join_ret1, join_ret2;	
	
	printf("we get into the main_thread.\n");	

	err1 = pthread_create(&tid1, NULL, thread_fun1, (void*)0);
	err2 = pthread_create(&tid2, NULL, thread_fun2, (void*)0);
	if(err1 != 0 || err2 != 0){
		printf("create new thread failed.\n");
		return 0;
	}
	
	join_ret1 = pthread_join(tid1, &rval1);
	join_ret2 = pthread_join(tid2, &rval2);
	printf("the join_ret1 is:%d\n", join_ret1);
	printf("the join_ret2 is:%d\n", join_ret2);
	

	sleep(1);
	printf("step out of main thread\n");
	return 0;
}
