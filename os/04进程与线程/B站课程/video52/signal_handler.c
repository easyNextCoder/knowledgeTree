/*
 *作用：验证sigaction函数，测试信号的发送和阻塞
 *作者：徐永康
 *日期：2019.11.22
 *说明：如果在thread_fun1和thread_fun2中都打开thread_sigmask
 *		则fun1和fun2则会阻塞主线程发来的SIGINT函数。
 *		如果任何一个线程中的阻塞函数被屏蔽则这个线程可以接收到信号
 *      接收到信号后，由于sigaction已经注册了SIGINT函数所以当线程
 * 		接收到SIGINT信号后会调用相关函数，而且一个SIGINT信号最终
 * 		只能对应一个函数，所以sigaction最后一次注册的谁，谁就最终
 * 		对应着SIGINT信号。
 */ 

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>

void sig_handler1(int arg){
	printf("thread1 get signal.\n");
	return ;
}

void sig_handler2(int arg){
	printf("thread2 get signal.\n");
	return ;
}

void* thread_fun1(void *arg){

	printf("new thread 1.\n");

	struct sigaction act;
	memset(&act, 0, sizeof(act));
	sigaddset(&act.sa_mask, SIGQUIT);
	act.sa_handler = sig_handler1;
	sigaction(SIGQUIT, &act, NULL);

	pthread_sigmask(SIG_BLOCK, &act.sa_mask, NULL);
	sleep(2);
	
}

void* thread_fun2(void *arg){
	
	printf("new thread 2.\n");

	struct sigaction act;
	memset(&act, 0, sizeof(act));
	sigaddset(&act.sa_mask, SIGQUIT);
	act.sa_handler = sig_handler2;
	sigaction(SIGQUIT, &act, NULL);

	pthread_sigmask(SIG_BLOCK, &act.sa_mask, NULL);
	sleep(2);
	
}

int main(){

	pthread_t tid1, tid2;
	int err;
	int s;
	void *rval;
	
	err = pthread_create(&tid1, NULL, thread_fun1, NULL);
	if(err != 0){
		printf("we can not create new pthread.\n");
		return 0;
	}

	err = pthread_create(&tid2, NULL, thread_fun2, NULL);
	if(err != 0){
		printf("we can not create new pthread.\n");
		return 0;
	}

	sleep(1);

	s = pthread_kill(tid1, SIGQUIT);
	if(s != 0){
		printf("send signal to thread1 failed.\n");
	}
	s = pthread_kill(tid2, SIGQUIT);
	if(s != 0){
		printf("send signal to thread2 failed.\n");
	}

	//两个join函数等待子线程运行完毕
	pthread_join(tid1, NULL);
	pthread_join(tid2, NULL);

	return 0;
}
