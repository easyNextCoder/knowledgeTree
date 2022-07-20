#include <iostream>
#include <thread>
#include <sstream>
#include <queue>
#include <string>
//linux header file
//#include <pthread>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
//personal headers
#include "MessageFormat.h"

using namespace std;
int SERVER_ID = -2;
//int self_client_id = -2;

int init_addr()
{
	
	return 0;
}

int self_client_id = -1;

void handle_recv(int sock)
{
	MessageBody ims;
	int rval = 1;
	while(rval>0)
	{
		rval = recv(sock, reinterpret_cast<char*>(&ims), sizeof(MessageBody), 0);
	if(ims.get_src_id() == SERVER_ID)
		{
			cout<<"Assigned ID is: "<<ims.get_des_id()<<endl;
			self_client_id = ims.get_des_id();
		}else{
			cout<<"Message from "<<ims.get_src_id()<<":"<<ims.get_message()<<endl;
		}
	}
}

int main()
{
	
	int err = 0;
	int sock = socket(AF_INET, SOCK_STREAM, 0);

	struct sockaddr_in addr;
	//socklen_t len = sizeof(struct sockaddr);
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = inet_addr("127.0.0.1");
	addr.sin_port = htons(3000);
A:
        int rval = connect(sock,  (const struct sockaddr*)(&addr), (socklen_t)sizeof(struct sockaddr));

	if(rval < 0)
	{
		cout<<"client:connect to server failed."<<endl;
		usleep(500000);
		goto A;
	}
	cout<<"client: get connected."<<endl;
	//new a thread;
	thread recv_handle_thread(handle_recv, sock);
	
	while(1)
	{
		int des_id = 0;
		string total_string, int_string, sendOutString;
		cout<<"Message to ";
		cin>>des_id;
		getline(cin, sendOutString);
		MessageBody sendOutMessage(self_client_id, des_id, sendOutString); 
		cout<<"client sendOutMessage:"<<sendOutMessage.get_message()<<endl;
		send(sock, reinterpret_cast<char*>(&sendOutMessage), sizeof(MessageBody), 0);
		usleep(500);
	}
	return 0;
}
