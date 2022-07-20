/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：运行这个server之后这个server最后支持10个client与之建立连接并发送最大100个字节的字符串
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

#define MAX_LISTEN_NUM 10
#define MAX_CLIENT_NUM 10

char buf[100];
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

void * thread_fun(void *arg){

    printf("I am thread_fun %d\n", (int)arg);
    int ret = 0;
    while(1){

        memset(buf, 0, 100);

        ret = read(client_request[(int)arg], buf, 100);
        if(ret <= 0){
            break;
        }else{
            if( is_equal(buf, "add", 3) ){

                pthread_mutex_lock(&mutex_common_id);
                common_id++;
                printf("after thread %d add operation , common_id is %d\n", (int)arg, common_id);
                pthread_mutex_unlock(&mutex_common_id);
            
            }
            printf("buf is %s, ret is %d\n", buf,ret);
        }
        
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
    pthread_t tid[MAX_CLIENT_NUM] ;
    while( client_request_cnt-- ){
        
        //等待客户端的请求
        client_request[client_request_cnt] = accept(sd, (struct sockaddr *)(&remote_ip), &remote_len);
        if(client_request[client_request_cnt] == -1){
            printf("accept error, errno is %d\n", errno);
            close(sd);
            return 0;
        }
        int err = pthread_create(&tid[client_request_cnt], NULL, thread_fun, (void*)client_request_cnt);
        
    }

    //检查是否还有线程活着如果有线程或者就汇入，等待其他线程运行完毕，再运行主线程即退出
    for(int i = 0; i<MAX_CLIENT_NUM; i++){
        int rval_pthread_kill = pthread_kill(tid[i], NULL);
        if(rval_pthread_kill == ESRCH){
            printf("thread %d is dead.\n", i);
        }else if(rval_pthread_kill == EINVAL){
            printf("signal is invalid.\n");
        }else{
            int err = pthread_join(tid[i], NULL);
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
