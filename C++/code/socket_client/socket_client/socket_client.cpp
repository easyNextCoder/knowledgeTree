// thread example
#include <iostream>       // std::cout
#include <thread>         // std::thread
#include <WINSOCK2.h>
#pragma comment(lib, "ws2_32.lib")
#pragma  warning(disable:4996) 
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
    init_addr(addr, AF_INET, inet_addr("127.0.0.1"), htons(3000));
A:
    int failed = connect(sk, (sockaddr*)&addr, sizeof(addr));
    cout << "after connect." << endl;
    if (failed < 0) {
        cout << "client:connect to server failed." << endl;
        goto A;
    }

    char send_buf[100] = "client: I am starting to connect to you.";
    char recv_buf[100];
    string get_string;
    while (1) {
        cin >> get_string;
        cin.ignore();
        for (int i = 0; i < get_string.size(); i++)
            send_buf[i] = get_string[i];
        send_buf[get_string.size()] = '\0';
        send(sk, send_buf, 100, 0);
        //recv(sk, recv_buf, 100, 0);

        cout << "client send out." << endl;
        Sleep(500);
    }



    return 0;
}