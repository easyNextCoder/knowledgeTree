/* 日期：2019/11/25
 * 作者：徐永康
 * 作用：
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

#define MAX_FRAME_LENGTH 100
struct sockaddr_in server_ip,  remote_ip;
int sd;
int assigned_id_by_server = -1;


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

char* get_des_id_from_send_command(char * input_command_copy, int * ptr_des_send_id){

    input_command_copy = input_command_copy + 4;
    int i = 0;
    while( *input_command_copy == ' '){
        input_command_copy++;
    }
    if( *input_command_copy == '\0'){
        ;
    }else{
        int handred_bit = 1;
        int decade_bit = 1;
        int unit_bit = 1;
        *ptr_des_send_id = (*input_command_copy - '0')*100 + (*(input_command_copy + 1) - '0')*10 + (*(input_command_copy + 2) - '0')*1;
        input_command_copy += 3;
    }
    return input_command_copy;
}

int load_command_data_to_frame_data(char * command_left, char * ptr_frame_data){
    int length = 0;
    char * command_left_copy = command_left;
    while(*command_left++ != '\0'){
        length++;
    }
    memcpy(ptr_frame_data, command_left_copy, length);

}

static Data_frame client_send_frame;
//client_send_frame 需要先init为小于0的数字
Data_frame * command_solve(char * input_command, int command_max_length){
     
    if(client_send_frame.is_connected_to_server <= 0){
        //并没有连接成功
        if( is_equal(input_command, "connect", 7) ){

            //准备去与服务器连接
            client_send_frame.data_type = CONNECT;
            memset(client_send_frame.data, sizeof(client_send_frame.data), 0);
            client_send_frame.is_solved_correctly = 1;
        }else{
            
            client_send_frame.is_solved_correctly = 0;
        }

    }else{
        //已经成功连接接下来进行send 和 recv 命令的解析
        if( is_equal(input_command, "send", 4) ){

            client_send_frame.is_solved_correctly = 1;
            client_send_frame.data_type = SEND;
            //解析发送到哪个客户端
            char * command_left;
            int des_send_id = 0;
            command_left = get_des_id_from_send_command(input_command, &des_send_id);
            client_send_frame.des_client_id = des_send_id;
            //装载要发送的数据到数据帧
            memset(client_send_frame.data, sizeof(client_send_frame.data), 0);
            load_command_data_to_frame_data(command_left, client_send_frame.data);

        }else{
            //并不是正确的命令
            client_send_frame.is_solved_correctly = 0;
        }
    }
        
    return &client_send_frame;

}

void * thread_write(void *arg){

    char *send_buf = (char*)malloc(sizeof(char)*MAX_FRAME_LENGTH) ;
    pthread_mutex_t socket_mutex;
    
    while( gets(send_buf) != EOF){
    
        //内部是  命令：发送内容  的循环
        command_solve(send_buf, MAX_FRAME_LENGTH);
        pthread_mutex_lock(&socket_mutex);
        //传输的时候使用互斥量进行加锁，用读写锁应该更好
        write(sd, &client_send_frame, sizeof(Data_frame));
        pthread_mutex_unlock(&socket_mutex);
        memset(send_buf, MAX_FRAME_LENGTH, 0);
    }
    free(send_buf);
    return NULL;
}

int read_from_server(){

    pthread_mutex_t socket_mutex;
    pthread_mutex_lock(&socket_mutex);
    read(sd, &client_send_frame, MAX_FRAME_LENGTH);
    pthread_mutex_unlock(&socket_mutex);

}

void * thread_read(void *arg){

    while(1){
        read_from_server();
        if(client_send_frame.data_type == CONNECT){    
            printf("connected to server successfully, assigned socket id is:%3.0d\n", client_send_frame.src_client_id);
        }else if(client_send_frame.data_type == SEND){
            printf("received data from socket id %3.0d : %s\n", client_send_frame.src_client_id, client_send_frame.data);
        }
    }
    return NULL;
}

int main(){

    
    pthread_t tid_read, tid_write;

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

    int err = connect(sd, (struct sockaddr *)(&server_ip), sizeof(struct sockaddr));
    if(err == -1){
        printf("connect error.\n");
        close(sd);
        return 0;
    }

    err = pthread_create(&tid_read, NULL, thread_read, NULL);
    if( err != 0 ){
        printf("thread_read create failed.\n");
        close(sd);
        return 0;
    }

    err = pthread_create(&tid_write, NULL, thread_write, NULL);
    if( err != 0 ){
        printf("thread_write create failed.\n");
        close(sd);
        return 0;
    }

    pthread_join(tid_write, NULL);
    pthread_join(tid_read, NULL);
    close(sd);
    return 0;
}
