#include <iostream>
#include <string>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/epoll.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <random>
#include <map>
#include <thread>
#include <assert.h>
#include <unistd.h>
#include "ServerMessage.h"
using namespace std;

map<int, int>id_socket;
map<int, int>socket_id;
/*
//use multithread!
void handle_client(int sock)
{
	static int id = 0;
	id++;
	cout<<"you are number: "<<id<<" to connect the server."<<endl;
	MessageBody ims;
	int rval = 1;
	while(rval > 0)
	{
		rval = recv(sock, reinterpret_cast<char*>(&ims), sizeof(MessageBody), 0);
		//cout<<ims._ms<<endl;
		ServerMessage sims(ims);
		
		send(id_socket[sims.get_des_id()], reinterpret_cast<char*>(&ims), sizeof(MessageBody), 0);
		cout<<"Received message:"<<sims.get_message()<<endl;
		cout<<"From client id: "<<sims.get_src_id()<<endl;
	}
}
*/
int main()
{
	int sock, epfd;
	struct epoll_event ev,ev1, events[20];
	epfd = epoll_create(256);
	sock = socket(AF_INET, SOCK_STREAM , 0);

	ev.data.fd = sock;
	ev.events = EPOLLIN|EPOLLET;
	//register event
	epoll_ctl(epfd, EPOLL_CTL_ADD, sock, &ev);

	struct sockaddr_in addr;
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = htonl(INADDR_ANY);
	addr.sin_port = htons(3000);
	bind(sock, (struct sockaddr*)&addr, sizeof(struct sockaddr));
	listen(sock, 10);

	int serConn[10];
	thread thread_con[10];
	vector<int> id_con;
	default_random_engine e;
	
	for(int i = 0; i<10; i++)
	{
		id_con.push_back(i+1);
	}

	int index = 0;
	int nfds = 0;
	MessageBody ims;
	MessageBody* s2c = new MessageBody(-2, id_con.back(), "");
	MessageBody* s2cBak = s2c;	
	for(;;)
	{
		
		nfds = epoll_wait(epfd, events, 10, 500);
		
		for(int i = 0; i<nfds; ++i)
		{
			if(events[i].data.fd == sock)
			{
				int tmpSock  = accept(sock, NULL, NULL);
				id_socket.insert(make_pair(id_con.back(), tmpSock));
				socket_id.insert(make_pair(tmpSock, id_con.back()));
				
				s2cBak->set_des_id(id_con.back());
				id_con.pop_back();
				s2c = s2cBak;
				ev.data.fd = tmpSock;
				ev.events = EPOLLOUT|EPOLLET;
				epoll_ctl(epfd, EPOLL_CTL_ADD, tmpSock, &ev);	
				
			}else if(events[i].events & EPOLLIN)
			{
				int sockfd = events[i].data.fd;
				
				int rval = recv(sockfd, reinterpret_cast<char*>(&ims), sizeof(MessageBody), 0);
				
				if(rval == 0)
				{
					assert(socket_id.count(sockfd) != 0);
					int tmpID = socket_id[sockfd];
					assert(id_socket.count(tmpID) != 0);
					socket_id.erase(sockfd);
					id_socket.erase(tmpID);
					id_con.push_back(tmpID);
					
					ev.events = EPOLLOUT;
					ev.data.fd = sockfd;
					epoll_ctl(epfd, EPOLL_CTL_DEL, sockfd, &ev);
					close(sockfd);
					continue;
				}
				//cout<<ims.get_des_id()<<endl;
				ev.events = EPOLLOUT;
				ev.data.fd = id_socket[ims.get_des_id()];
				s2c = &ims;
				epoll_ctl(epfd, EPOLL_CTL_MOD, id_socket[ims.get_des_id()], &ev);
				ev1.events = EPOLLIN;
				ev1.data.fd = sockfd;
				epoll_ctl(epfd, EPOLL_CTL_MOD, sockfd, &ev1);			
				//return -1;
			}else if(events[i].events & EPOLLOUT)
			{
				//MessageBody * pims = events[i].data.ptr;
				send(events[i].data.fd, reinterpret_cast<char*>(s2c), sizeof(MessageBody), 0);
				cout<<"send to"<<s2c->get_des_id()<<endl;
			
				ev.events = EPOLLIN;
				epoll_ctl(epfd, EPOLL_CTL_MOD, events[i].data.fd,&ev); 
				
			}
		}
	}

	
	return 0;
}
