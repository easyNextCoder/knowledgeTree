/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：作为聊天程序的服务器来使用
 */ 
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>
#include <errno.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>

#include "data_frame.h"
#include "errors.h"

#define MAX_LISTEN_NUM 10
#define MAX_CLIENT_NUM 10
#define MAX_FRAME_LENGTH 100


int client_request[MAX_CLIENT_NUM];
int common_id = 0;
pthread_mutex_t mutex_common_id;

int is_equal(char * s1, char *s2, int length){
    while(s2[length-1] == s1[length-1]){
        if(length - 1 == 0){
            return 1;
        }else{
            length--;
        }
    }
    return 0;
}

int write_backto_client(const Data_frame *send_out_frame){
    
    pthread_mutex_t socket_mutex;
    switch(send_out_frame->data_type){
        case(CONNECT):
            //在已经有互斥量锁的区间中不能再加锁
            //pthread_mutex_lock(&socket_mutex);
            write(client_request[send_out_frame->src_client_id], send_out_frame, sizeof(Data_frame));
            //pthread_mutex_unlock(&socket_mutex);
            break;
        case(SEND):
            write(client_request[send_out_frame->des_client_id], send_out_frame, sizeof(Data_frame));
            break;
        default:
            return -SOLVE_ERR;
    }   

}

int frame_handler(Data_frame * received_frame, int assigned_socket_id){
    
    //对不同的请求分别进行不同的处理
    if( !received_frame->is_solved_correctly ){
        return -SOLVE_ERR;
    }
    switch(received_frame->data_type){
        case(CONNECT):
            received_frame->is_connected_to_server = 1;
            received_frame->src_client_id = assigned_socket_id;
            write_backto_client(received_frame);
            break;
        case(SEND):
            write_backto_client(received_frame);
            break;
        default:
            return -SOLVE_ERR;
    }
}

void * thread_read(void *arg){

    int ret = 0;
    pthread_mutex_t socket_mutex;
    char buf[MAX_FRAME_LENGTH] = {0};
    while(1){
        printf("start thread %d\n", arg);
        int err = pthread_mutex_lock(&socket_mutex);
        if(err != 0){
            printf("ERROR:%s\n", errors[MUTEX_LOCK_ERR]);
        }
        //这里当client没有发送的时候，整个程序的运行会暂停在read整个函数这里等待client的发送
        //当client与server断开连接后，整个程序会快速的运行。在client中屏蔽互斥量也是这样的效果
        //可见也并不是互斥量起的作用。
        ret = read(client_request[(int)arg], buf, MAX_FRAME_LENGTH);
        Data_frame * received_frame = (Data_frame*)buf;
        pthread_mutex_unlock(&socket_mutex);
        //已经读上来数据帧，接下来对数据帧进行处理
        int rval = frame_handler(received_frame, (int)arg);
        if(rval < 0){
            printf("ERROR:%s\n", errors[-rval]);
        }

        memset(buf, MAX_FRAME_LENGTH, 0);
    }
    close(client_request[(int)arg]);
    return NULL;
}

int main(){

    struct sockaddr_in server_ip,  remote_ip;
    int sd;

    //创建socket
    sd = socket(AF_INET, SOCK_STREAM, 0);
    if(sd == -1){
        printf("create socket failed, errno is %d\n", errno);
        return 0;
    }

    //设置IP地址和端口
    server_ip.sin_family = AF_INET;
    server_ip.sin_port = htons(1234);
    server_ip.sin_addr.s_addr = htonl(INADDR_ANY);
    memset(server_ip.sin_zero, 0, 8);

    //绑定IP地址和端口到socket
    int err = bind(sd, (struct sockaddr*)(&server_ip), sizeof(struct sockaddr));
    if(err == -1){
        printf("bind error. errno is %d\n", errno);
        close(sd);
        return 0;
    }

    //设置服务器的最大连接数
    err = listen(sd, MAX_LISTEN_NUM);
    if(err == -1){
        printf("listen error, errno is:%d\n", errno);
        close(sd);
        return 0;
    }

    //获取客户端IP地址的长度
    int remote_len = sizeof(struct sockaddr);
    int client_request_cnt = MAX_CLIENT_NUM;
    pthread_t tid_write[MAX_CLIENT_NUM] ;
    while( client_request_cnt-- ){
        
        //等待客户端的请求
        client_request[client_request_cnt] = accept(sd, (struct sockaddr *)(&remote_ip), &remote_len);
        if(client_request[client_request_cnt] == -1){
            printf("accept error, errno is %d\n", errno);
            close(sd);
            return 0;
        }
        int err = pthread_create(&tid_write[client_request_cnt], NULL, thread_read, (void*)client_request_cnt);
        
    }

    //检查是否还有线程活着如果有线程或者就汇入，等待其他线程运行完毕，再运行主线程即退出
    for(int i = 0; i<MAX_CLIENT_NUM; i++){
        int rval_pthread_kill = pthread_kill(tid_write[i], NULL);
        if(rval_pthread_kill == ESRCH){
            printf("thread %d is dead.\n", i);
        }else if(rval_pthread_kill == EINVAL){
            printf("signal is invalid.\n");
        }else{
            int err = pthread_join(tid_write[i], NULL);
            if(err != 0){
                printf("join failed.\n");
            }else{
                printf("main thread join pthread %d success.\n", i);
            }
        }
    }
    
    close(sd);
    return 0;
}
