/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：测试线程与fork函数配合使用
 * 
 */ 

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

void * thread_fun1(void *arg){

    sleep(1);
    pid_t pid;
    pid = fork();
    if(pid == 0){
        //这里是子进程，由于父进程已经对互斥量进行加锁，所以子进程无法再次加锁
        //从而就是卡死在这里
        pthread_mutex_lock(&mutex);
        printf("child\n");
        pthread_mutex_unlock(&mutex);
    }
    if(pid > 0){
        pthread_mutex_lock(&mutex);
        printf("parent.\n");
        pthread_mutex_unlock(&mutex);
    }

    return NULL;
}

int main(){

    pthread_t thread1, thread2;
    int err1, err2;

    err1 = pthread_create(&thread1, NULL, thread_fun1, 0);    
    if(err1 != 0){
        printf("create thread1 failed.\n");
    }

    pthread_mutex_lock(&mutex);
    sleep(2);
    pthread_mutex_unlock(&mutex);
    printf("main.\n");
    //这句话是将主线程汇入到线程thread1，就是就算main的主体代码运行完也要等待
    //线程thread1运行完整体才算结束，没有这句话只能打印出main
    pthread_join(thread1, NULL);
/*
    err1 = pthread_join(thread1, &rval);
    if(err1 != 0){
        printf("join pthread1 success.\n");
    }else{
        printf("join thread1 failed.\n");
    }
*/

    return 0;
}
