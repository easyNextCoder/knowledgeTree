/* 日期：2019/11/24
 * 作者：徐永康
 * 作用：验证一次性初始化问题
 * 结果：thread_fun1先于thread_fun2运行，最终thread_init只运行了一次
 */ 
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

pthread_once_t once = PTHREAD_ONCE_INIT;
pthread_t tid;

void thread_init(){
    printf("I am in thread 0x%x\n", tid);
}

void * thread_fun1(void *arg){
    tid = pthread_self();
    printf("I am thread 0x%x\n", tid);
    printf("once is %d\n", once);
    pthread_once(&once, thread_init);
    printf("once is %d\n",once);
    return NULL;
}

void *thread_fun2(void * arg){
    sleep(2);
    tid = pthread_self();
    printf("I am thread 0x%x\n",tid);
    pthread_once(&once, thread_init);

    return NULL;
}

int main(){

    pthread_t thread1, thread2;
    void *rval;

    pthread_create(&thread1, NULL, thread_fun1, 0);
    pthread_create(&thread2, NULL, thread_fun2, 0);

    pthread_join(thread1, &rval);
    pthread_join(thread2, &rval);

    return 0;
}
