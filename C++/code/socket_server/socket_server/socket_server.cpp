// thread example
#include <iostream>       // std::cout
#include <thread>         // std::thread
#include <WINSOCK2.h>
#include <stdlib.h>
#include <vector>
#include <algorithm>
#pragma comment(lib, "ws2_32.lib")
using namespace std;

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

void handle_client(SOCKET serConn) {
    static int id = 0;
    id++;
    cout << "you are number: " << id << " to connect the server." << endl;
    char recv_buf[100];
    while (1) {

        recv(serConn, recv_buf, 100, 0);
        /*
            显示和处理错误信息
            WSAGetLastError();
            wprintf(L"accept failed with error: %ld\n", WSAGetLastError());
         */
        {
            cout << recv_buf ;
        }
        
    }
    
}
int sock_con[100];
int main() {
    WORD request = MAKEWORD(1, 1);
    WSADATA wsa_data;
    int err = 0;
    err = WSAStartup(request, &wsa_data);
    //首先启动网络，request是对应的版本号
    if (!err) {
        cout << "DLL版本协商成功。" << endl;
    }

    SOCKET sk = socket(AF_INET, SOCK_STREAM, 0);
    //IP4协议，SOCK流
    SOCKADDR_IN addr;
    init_addr(addr, AF_INET, htonl(INADDR_ANY), htons(3000));

    bind(sk, (sockaddr*)&addr, sizeof(addr));

    listen(sk, 10);

    SOCKADDR_IN clientsocket;
    int len = sizeof(SOCKADDR);
    SOCKET serConn[20];
    thread thread_con[10];
    int index = 0;
    while (1) {

        //等待客户端的连接
        cout << "waitting for connect." << endl;
        serConn[index] = accept(sk, NULL, NULL);

        if (serConn[index] == INVALID_SOCKET) {
            wprintf(L"accept failed with error: %ld\n", WSAGetLastError());
            closesocket(serConn[index]);
            WSACleanup();
            return 1;
        }
        else
            wprintf(L"Client connected.\n");

        thread_con[index] = thread(handle_client, serConn[index]);
        index++;
        
    }

    return 0;
}
