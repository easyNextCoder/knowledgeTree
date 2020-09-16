#ifndef __DATA_FRAME__
#define __DATA_FRAME__

#define DATA_LENGTH 50
typedef enum {
	CONNECT = 1,//要进行连接
	SEND,		//要进行发送
}Data_type;
typedef struct  {
	int is_solved_correctly;
	int is_connected_to_server;
	Data_type data_type;
	int src_client_id;
	int des_client_id;
	char data[DATA_LENGTH];
} Data_frame;


#endif