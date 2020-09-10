/* 日期：2019/11/23
 * 作者：徐永康
 * 作用：生产者与消费者问题
 * 结果：程序运行的结果就是起初生产的快消费的满，等仓库满之后就是消费一个生产一个。
 */ 

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define BUFFER_SIZE    5            //产品库存大小
#define PRODUCT_CNT     30          //产品生产总数

struct product_cons{
    int buffer[BUFFER_SIZE];        //生产产品值
    pthread_mutex_t lock;           //互斥锁volatile int
    int readpos, writepos;          //读写位置
    pthread_cond_t notempty;        //条件变量，非空
    pthread_cond_t notfull;
};

struct product_cons buffer;

void init(struct product_cons *p){

    pthread_mutex_init(&p->lock, NULL);         //互斥锁
    pthread_cond_init(&p->notempty, NULL);      //条件变量
    pthread_cond_init(&p->notfull, NULL);       //条件变量
    p->readpos = 0;                             //读写位置
    p->writepos = 0;

}
void finish(struct product_cons *p){
    pthread_mutex_destroy(&p->lock);            //互斥锁
    pthread_cond_destroy(&p->notempty);         //条件变量
    pthread_cond_destroy(&p->notfull);          //读写位置
    p->readpos = 0;
    p->writepos = 0;
}

int get(struct product_cons *p){
    int data;

    pthread_mutex_lock(&p->lock);

    if(p->readpos == p->writepos){
        printf("consumer wait for not empty.\n");
        pthread_cond_wait(&p->notempty, &p->lock);
    }

    data = p->buffer[p->readpos++];
    if(p->readpos >= BUFFER_SIZE)
        p->readpos = 0;
    
    pthread_cond_signal(&p->notfull);
    pthread_mutex_unlock(&p->lock);
    return data;

}
int put(struct product_cons *p, int input_data){
    
    pthread_mutex_lock(&p->lock);
   
    if( (p->writepos + 1)%BUFFER_SIZE == p->readpos ){            //此时仓库已经满了
        printf("producer wait for not full.\n");       //等待仓库不满的时候才能进行put操作
        pthread_cond_wait(&p->notfull, &p->lock);
    }

    p->buffer[p->writepos++] = input_data;
    if(p->writepos >= BUFFER_SIZE)
        p->writepos = 0;

    pthread_cond_signal(&p->notempty);
    pthread_mutex_unlock(&p->lock);
}

void * produce(void *data){     //子线程，生产

    int n;
    for(n = 1; n <= 50; ++n){
        sleep(1);
        printf("put the %d product ...\n", n);
        put(&buffer, n);
        printf("put the %d product success.\n", n);
    }
    printf("producer stopped.\n");

    return NULL;
}

void * consumer(void *data){
    static int cnt = 0;
    int num;
    while(1){
        sleep(2);
        printf("get product ...\n");
        num = get(&buffer);
        printf("get the %d product success.\n", num);
        if( ++cnt == PRODUCT_CNT)
            break;
    }

    printf("consumer stopped.\n");
    return NULL;
}

int main(){
    pthread_t thread_consumer, thread_producer;
    void *rval;

    init(&buffer);

    pthread_create(&thread_consumer, NULL, consumer, 0);
    pthread_create(&thread_producer, NULL, produce, 0);

    pthread_join(thread_consumer, &rval);
    pthread_join(thread_producer, &rval);

    finish(&buffer);

    return 0;

    return 0;
}