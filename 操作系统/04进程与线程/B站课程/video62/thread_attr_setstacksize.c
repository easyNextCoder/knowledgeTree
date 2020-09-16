/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：修改线程默认栈的大小
 * 结果：
 */ 
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>

/*创建可分离的线程属性*/
    pthread_attr_t thread1_attr;

void * thread_fun1(void *arg){

    printf("I am thread fun1\n");
    size_t stack_size;
#ifdef _POSIX_THREAD_ATTR_STACKSIZE
    pthread_attr_getstacksize(&thread1_attr, &stack_size);
#endif
    printf("the thread1 stack_size is:%d\n", stack_size);
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

    
    //初始化属性
    pthread_attr_init(&thread1_attr);
    //设置为可分离的状态
    pthread_attr_setdetachstate(&thread1_attr, PTHREAD_CREATE_JOINABLE);
    
#ifdef _POSIX_THREAD_ATTR_STACKSIZE
    pthread_attr_setstacksize(&thread1_attr, PTHREAD_STACK_MIN);
    //默认情况下堆栈的大小是10M，当设置的栈的大小小于PTHREAD_STACK_MIN时
    //会出现设置不成功，不成功之后线程栈的仍然设置为默认的大小
#endif


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
