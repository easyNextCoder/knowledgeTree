#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void *thread_fun(void *arg){
    int stateval;
    int typeval;
    stateval = pthread_setcancelstate(PTHREAD_CANCEL_DISABLE, NULL);
    if(stateval != 0){
        printf("set cancel state failed.\n");
    }
    printf("I am new thread.\n");
    sleep(4);

    printf("about to cancel.\n");
    stateval = pthread_setcancelstate(PTHREAD_CANCEL_ENABLE, NULL);
    if(stateval != 0){
        printf("set cancel state failed.\n");
    }
    typeval = pthread_setcanceltype(PTHREAD_CANCEL_ASYNCHRONOUS, NULL);
    if(typeval != 0){
        printf("set cancel type failed.\n");
    }
    printf("first cancel point.\n");
    printf("second cancel point.\n");
    
    return (void*)20;
}

int main(){
    pthread_t tid;
    int err;
    err = pthread_create(&tid, NULL, thread_fun, NULL);
    if(err != 0){
        printf("can not create new thread.\n");
    }
    sleep(2);
    int cancel_ret = pthread_cancel(tid);
    if(cancel_ret != 0){
        printf("can not cancel the new thread.\n");
    }
    void * rval;
    int join_ret = pthread_join(tid, &rval);

    printf("new thread exit code is:%d\n", (int*)rval);

    return 0;
}