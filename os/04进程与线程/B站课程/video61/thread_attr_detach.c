#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

/*
 *PTHREAD_CREATE_DETACHED分离的
 * PTHREAD_CREATE_JOINABLE非分离的，可连接的
 * int pthread_attr_setdetachstate(pthread_attr_t *attr, int detachstate);
 * int pthread_attr_getdetachstate(pthread_attr_t *attr, int *detachstate);
 * 使用pthread_attr_getdetachstate可以获得线程的分离状态属性
 * 
 * 设置线程分离属性的步骤
 * 1.定义线程属性变量pthread_attr_t attr
 * 2.初始化attr,pthread_attr_init(&attr)
 * 3.设置线程为分离或非分离pthread_attr_setdetachstate(&attr, detachstate)
 * 4.创建线程pthread_create(&tid, &attr, thread_fun, NULL)
 * 所有的系统都会支持线程的分离状态属性
 */


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

void * thread_fun1(void *arg){

    printf("I am thread fun1\n");

    return NULL;
}

void *thread_fun2(void * arg){

    printf("I am thread fun2\n");

    return NULL;
}

int main(){

    pthread_t thread1, thread2;
    void *rval;
    int err1, err2;

    /*创建可分离的线程属性*/
    pthread_attr_t thread1_attr;
    //初始化属性
    pthread_attr_init(&thread1_attr);
    //设置为可分离的状态
    pthread_attr_setdetachstate(&thread1_attr, PTHREAD_CREATE_DETACHED);

    err1 = pthread_create(&thread1, &thread1_attr, thread_fun1, 0);
    
    if(err1 != 0){
        printf("create thread1 failed.\n");
    }

    err2 = pthread_create(&thread2, NULL, thread_fun2, 0);

    if(err2 != 0){
        printf("create thread2 failed.\n");
    }

    err1 = pthread_join(thread1, &rval);
    if(err1 != 0){
        printf("join pthread1 success.\n");
    }else{
        printf("join thread1 failed.\n");
    }
    
    err2 = pthread_join(thread2, &rval);
    if(err2 != 0){
        printf("join pthread2 success.\n");
    }else{
        printf("join thread2 failed.\n");
    }

    //注意属性要记得去销毁
    pthread_attr_destroy(&thread1_attr);

    return 0;
}
