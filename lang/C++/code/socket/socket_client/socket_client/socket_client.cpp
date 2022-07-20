// thread example
#include <iostream>       // std::cout
#include <thread>         // std::thread
#include <sstream>
#include <WINSOCK2.h>
#include <queue>

#include "socket_client.h"

#pragma comment(lib, "ws2_32.lib")
#pragma  warning(disable:4996) 
using namespace std;

int self_client_id = -2;



class A {
public:
    A();
};
class B {
public:
    B(int a);

};
void dosomething(B obj) {
    ;
}
/*
using namespace std;

void foo()
{
    // do stuff...
    while (1) {
        cout << "doing foo() function." << endl;
    }
}

void bar(int x)
{
    // do stuff...
    while (1) {
        cout << "doing bar() function." << endl;
    }
}

int main()
{
    std::thread first(foo);     // spawn new thread that calls foo()
    std::thread second(bar, 0);  // spawn new thread that calls bar(0)

    std::cout << "main, foo and bar now execute concurrently...\n";

    // synchronize threads:
    first.join();                // pauses until first finishes
    second.join();               // pauses until second finishes

    std::cout << "foo and bar completed.\n";

    return 0;
}
*/
int init_addr(SOCKADDR_IN& addr, USHORT ip_version, ULONG ip_addr, USHORT port) {
    addr.sin_family = ip_version;
    addr.sin_addr.S_un.S_addr = ip_addr;//ip地址
    addr.sin_port = port;//绑定端口

    return 1;
}


void handle_recv(SOCKET sk){
	message_body ims;
	while(1){
		recv(sk, reinterpret_cast<char*>(&ims), sizeof(message_body), 0);
		message_base oms(ims);
		
        if (oms.get_src_id() == SERVER_ID) {
            cout << "Assigned ID is: " << oms.get_des_id() << endl;

            self_client_id = oms.get_des_id();
        }
        else {
            cout << "Message from  " << oms.get_src_id() <<" :"<<oms.get_message() << endl;
        }
        
	}
}



int main() {
    priority_queue<int> mq;
    
    WORD request = MAKEWORD(1, 1);
    WSADATA wsa_data;
    int err = 0;
    err = WSAStartup(request, &wsa_data);
    //首先启动网络，request是对应的版本号
    if (!err) {
        cout << "DLL版本协商成功。" << endl;
    }
    dosomething(14);
    SOCKET sk = socket(AF_INET, SOCK_STREAM, 0);
    //IP4协议，SOCK流
    SOCKADDR_IN addr;
    init_addr(addr, AF_INET, inet_addr("127.0.0.1"), htons(3000));
A:
    int failed = connect(sk, (sockaddr*)&addr, sizeof(addr));
    cout << "after connect." << endl;
    if (failed < 0) {
        cout << "client:connect to server failed." << endl;
        goto A;
    }

    
	thread recv_handle_thread(handle_recv, sk);
	
	while (1) {
        int des_id = 0;
        string  total_string, int_string, send_out_string;
        cout << "Message to ";
        cin >> des_id;
        getline(cin, send_out_string);
        //stringstream stringline(total_string);
        //stringline >>int_string;
        
        
        //des_id = stoi(int_string, NULL);

        message_body sendout_message(self_client_id, des_id, send_out_string);

        send(sk, reinterpret_cast<char*>(&sendout_message), sizeof(message_body), 0);

        Sleep(500);
    }

    return 0;
}


