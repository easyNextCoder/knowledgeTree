// thread example
#include <iostream>       // std::cout
#include <thread>         // std::thread
#include <WINSOCK2.h>
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
int init_addr(SOCKADDR_IN& addr, ULONG ip_version, ULONG ip_addr, USHORT port) {
    addr.sin_family = AF_INET;
    addr.sin_addr.S_un.S_addr = htonl(INADDR_ANY);//ip地址
    addr.sin_port = htons(3000);//绑定端口

    return 1;
}
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
    SOCKET serConn;
    //等待客户端的连接
    serConn = accept(sk, (SOCKADDR*)&clientsocket, &len);

    char send_buf[100] = "server: you have accessed server.";
    char recv_buf[100];
    while (1) {

        send(serConn, send_buf, 100, 0);
        recv(serConn, recv_buf, 100, 0);

        cout << recv_buf << endl;
    }



    return 0;
}