/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：测试线程的私有数据
 * 
 */ 

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>

pthread_key_t key;

void * thread_fun1(void *arg){

    printf("I am thread fun1\n");
    int a = 1;
    //将键key与变量a关联
    pthread_setspecific(key, (void*)a);
    sleep(2);
    printf("thread 1 key->data is %d\n", pthread_getspecific(key));
    return NULL;
}

void *thread_fun2(void * arg){
    sleep(1);
    printf("I am thread fun2\n");
    int a = 2;
    pthread_setspecific(key, (void*)a);
    printf("thread 2 key->data is %d\n", pthread_getspecific(key));
    return NULL;
}

int main(){

    pthread_t thread1, thread2;
    void *rval;
    int err1, err2;

    //在这里先设置一个键
    pthread_key_create(&key, NULL);


    err1 = pthread_create(&thread1, NULL, thread_fun1, 0);
    
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

    //最后销毁这个键
    pthread_key_delete(&key);

    return 0;
}
