#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
int n = 10;

void* thread_fun1(void *arg){
	int random_number = -1;
	while(1){

		srand(time(NULL));
		random_number = rand()%n+1;
		printf("the random_number1 is:%d\n", random_number);
		if(random_number == n-1){
			printf("exit from thread_fun1.\n");
			exit(1);
		}
	}
}
void* thread_fun2(void *arg){
	int random_number = -1;
	while(1){

		srand(time(NULL));
		random_number = rand()%n+1;
		printf("the random_number2 is:%d\n", random_number);
		if(random_number == n-1){
			printf("exit from thread_fun2.\n");
			exit(2);
		}
	}
	
}
void* thread_fun3(void *arg){
	int random_number = -1;
	while(1){

		srand(time(NULL));
		random_number = rand()%n+1;
		printf("the random_number3 is:%d\n", random_number);
		if(random_number == n-1){
			printf("exit from thread_fun3.\n");
			exit(3);
		}
	}
}
void* thread_fun4(void *arg){
	int random_number = -1;
	while(1){

		srand(time(NULL));
		random_number = rand()%n+1;
		printf("the random_number4 is:%d\n", random_number);
		if(random_number == n-1){
			printf("exit from thread_fun4.\n");
			exit(4);
		}
	}
}




int main(int argc, char* argv[]){
	
	int err;
	pthread_t tid;
	
	srand(time(NULL));
	printf("the rand number is:%d\n",rand()%4+1);
		
	err = pthread_create(&tid, NULL, thread_fun1, (void*)argv[1]);
	if(err != 0){
		printf("create new thread1 failed.\n");
		return 0;
	}
	err = pthread_create(&tid, NULL, thread_fun2, (void*)argv[1]);
	if(err != 0){
		printf("create new thread2 failed.\n");
		return 0;
	}
	err = pthread_create(&tid, NULL, thread_fun3, (void*)argv[1]);
	if(err != 0){
		printf("create new thread3 failed.\n");
		return 0;
	}
	err = pthread_create(&tid, NULL, thread_fun4, (void*)argv[1]);
	if(err != 0){
		printf("create new thread4 failed.\n");
		return 0;
	}




	while(1)	
	sleep(1);
	printf("main thread\n");
	
	return 0;
}
