/*
 *作用：测试线程的读写锁操作
 *作者: 徐永康
 *日期：2019.11.23
 */

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int num = 0;
pthread_rwlock_t rwlock;

/*
//当两个函数都加上读的锁时候，两个线程是可以交替执行
void *thread_fun1(void *arg){

    pthread_rwlock_rdlock(&rwlock);
    printf("I am thread1:%d\n", num);
    sleep(3);
    printf("thread1 finish.\n");
    pthread_rwlock_unlock(&rwlock);

    return (void*)1;
}

void *thread_fun2(void *arg){

    pthread_rwlock_rdlock(&rwlock);
    printf("I am thread2:%d\n", num);
    sleep(3);
    printf("thread2 finish.\n");
    pthread_rwlock_unlock(&rwlock);
    
    return (void*)2;
}
*/
/*
//当两个线程都是写锁时只能等待着排队执行
void *thread_fun1(void *arg){

    pthread_rwlock_wrlock(&rwlock);
    printf("I am thread1:%d\n", num);
    sleep(3);
    printf("thread1 finish.\n");
    pthread_rwlock_unlock(&rwlock);

    return (void*)1;
}

void *thread_fun2(void *arg){

    pthread_rwlock_wrlock(&rwlock);
    printf("I am thread2:%d\n", num);
    sleep(3);
    printf("thread2 finish.\n");
    pthread_rwlock_unlock(&rwlock);
    
    return (void*)2;
}
*/
//当一个线程读，另一个线程要写的时候，为了内容的完整性的更改，也只能排队运行
void *thread_fun1(void *arg){

    pthread_rwlock_rdlock(&rwlock);
    printf("I am thread1:%d\n", num);
    sleep(3);
    printf("thread1 finish.\n");
    pthread_rwlock_unlock(&rwlock);

    return (void*)1;
}

void *thread_fun2(void *arg){

    pthread_rwlock_wrlock(&rwlock);
    printf("I am thread2:%d\n", num);
    sleep(3);
    printf("thread2 finish.\n");
    pthread_rwlock_unlock(&rwlock);
    
    return (void*)2;
}

int main(){

    pthread_t tid1, tid2;
    int err1 = 0;

    pthread_rwlock_init(&rwlock, NULL);

    err1 = pthread_create(&tid1, NULL, thread_fun1, NULL);
    if(err1 != 0){
        printf("create thread1 failed.\n");
        return 0;
    }
    
    err1 = pthread_create(&tid2, NULL, thread_fun2, NULL);
    if(err1 != 0){
        printf("create thread2 failed.\n");
        return 0;
    }

    pthread_join(tid1, NULL);
    pthread_join(tid2, NULL);

    printf("main thread finished.\n");
    pthread_rwlock_destroy(&rwlock);
    return 0;
}