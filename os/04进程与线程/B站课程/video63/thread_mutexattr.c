/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：互斥量属性的测试,设置互斥量的属性为共享的，可以将将互斥量同步到多个进程使用
 * 编译命令：gcc thread_mutexattr.c -lpthread -lrt
 */ 
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>
#include <sys/shm.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>


//下面这个程序，通过设置互斥量的属性为多进程共享的，
//那么就可以在多个进程之间进行互斥量的共享
int main(){

    char * shm = "myshm";
    char * shm1 = "myshm1";
    int shm_id, shm_id1;
    char *buf;
    pid_t pid;

    pthread_mutex_t *mutex;
    //这是一个属性的数据
    pthread_mutexattr_t mutexattr;

    //打开共享内存
    shm_id1 = shm_open(shm1, O_RDWR|O_CREAT, 0644);
    //调整共享内存大小
    ftruncate(shm_id1, 100);
    //映射共享内存，MAP_SHARED属性表明，对共享内存的任何修改都会影响其他进程
    mutex = (char*)mmap(NULL, 100, PROT_READ|PROT_WRITE, MAP_SHARED, shm_id1, 0);

    pthread_mutexattr_init(&mutexattr);
#ifdef _POSIX_THREAD_PROCESS_SHARED
    pthread_mutexattr_setpshared(&mutexattr, PTHREAD_PROCESS_SHARED);
                                             //PTHREAD_PROCESS_PRIVATE
#endif
    pthread_mutex_init(mutex, &mutexattr);

    pid = fork();
    if(pid == 0){
        //休眠1s，让父进程先运行
        sleep(1);
        printf("I am child process.\n");

        pthread_mutex_lock(mutex);
        //将共享内存内容修改未hello
        memcpy(buf, "hello", 6);
        printf("child buf is:%s\n", buf);
        pthread_mutex_unlock(mutex);
    }else if(pid > 0){
        printf("I am parent process.\n");

        pthread_mutex_lock(mutex);
        //修改共享内存到内容，改为world
        memcpy(buf, "world", 6);
        sleep(3);
        printf("parent buf is:%s\n", buf);
        pthread_mutex_unlock(mutex);
    }

    pthread_mutexattr_destroy(&mutexattr);
    pthread_mutex_destroy(mutex);

    //接触映射
    munmap(buf, 100);
    //消除共享内存
    shm_unlink(shm);
    return 0;
}
