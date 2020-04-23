
#include <iostream>       // std::cout
#include <thread>         // std::thread
#include <WINSOCK2.h>
#include <stdlib.h>
#include <vector>
#include <algorithm>
#include <random>
#include <map>

#include "socket_server.h"
#pragma comment(lib, "ws2_32.lib")

using namespace std;


map<int, SOCKET>id_socket;


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
    message_body ims;
    while (1) {

        recv(serConn, reinterpret_cast<char*>(&ims), sizeof(message_body), 0);
        
        server_message sims(ims);

        send(id_socket[sims.get_des_id()], reinterpret_cast<char*>(&ims), sizeof(message_body), 0);

        cout << "Received message." << endl;
        cout << "From client: " << sims.get_src_id()<< endl;
        cout << "Message is: " << sims.get_message() << endl;
    }
    
}

int main() {

    WORD request = MAKEWORD(1, 1);
    WSADATA wsa_data;
    int err = 0;
    err = WSAStartup(request, &wsa_data);
    //首先启动网络，request是对应的版本号
    if (!err) {
        output({ {"server information","DLL版本协商成功。"} });
    }

    SOCKET sk = socket(AF_INET, SOCK_STREAM, 0);
    //IP4协议，SOCK流
    SOCKADDR_IN addr;
    init_addr(addr, AF_INET, htonl(INADDR_ANY), htons(3000));

    bind(sk, (sockaddr*)&addr, sizeof(addr));

    listen(sk, 10);

    
    SOCKET serConn[MAX_THREAD_NUM];
    thread thread_con[MAX_THREAD_NUM];
    vector<int>id_con;
    default_random_engine e;
    for (int i = 0; i < MAX_THREAD_NUM; i++) {
        id_con.push_back(i + 1);
    }
    int index = 0;
    output({ {"server information","等待用户的连接。"} });
    
    
    while (1) {
        
        //等待客户端的连接
        serConn[index] = accept(sk, NULL, NULL);

        if (serConn[index] == INVALID_SOCKET) {
            wprintf(L"accept failed with error: %ld\n", WSAGetLastError());
            closesocket(serConn[index]);
            WSACleanup();
            return 1;
        }
        else {
            output({ {"server information","Client connected."} });
        }
        id_socket.insert(make_pair(id_con.back(),serConn[index]));
        send(serConn[index], reinterpret_cast<char*>(&message_body(SERVER_ID, id_con.back(), "ID assigned to you.")), sizeof(message_body),0);
        id_con.pop_back();
        thread_con[index] = thread(handle_client, serConn[index]);
        index++;
        
    }

    return 0;
}

        /*
            显示和处理错误信息
            WSAGetLastError();
            wprintf(L"accept failed with error: %ld\n", WSAGetLastError());
         */

//Ctrl + k ctrl + c
//ctrl + k ctrl + u
//#include <iostream>
//#include <vector>
//#include <math.h>
//#include <algorithm>
//#include <set>
//
//using namespace std;
//
//
//int  recal_knife(vector<pair<double, int> >& tmp, int knife) {
//
//    sort(tmp.begin(), tmp.end(), [](pair<double, int>& item, pair<double, int>& item2) {
//        return item.first < item2.first;
//        });
//    for (auto item : tmp) {
//        double knifep = knife;
//        if (knifep >= item.first) {
//            knife += item.second;
//        }
//    }
//    return knife;
//}
//int cal_length(vector< vector<int> >& con, int x, int y, int depth, int knife) {
//
//    int left_up_px = x - depth >= 0 ? x - depth : 0;
//    int left_up_py = y - depth >= 0 ? y - depth :0;
//    int right_dn_px = x + depth < con.size() ? x + depth : con.size() - 1;
//    int right_dn_py = y + depth < con.size() ? y + depth : con.size() - 1;
//    //上横行
//    vector<pair<double, int>> tmp;
//    for (int iy = left_up_py; iy < right_dn_py; iy++) {
//        double dis = sqrt(double((x - left_up_px) * (x - left_up_px)) +
//            double((y - iy) * (y - iy)));
//            int thing = con[left_up_px][iy];
//        if (thing == 0) {
//            continue;
//        }
//        else {
//            tmp.push_back(make_pair(dis, thing));
//        }
//        con[left_up_px][iy] = 0;
//    }
//    //右竖行 
//    for (int ix = left_up_px; ix < right_dn_px; ix++) {
//        double dis = sqrt(double((x - ix) * (x - ix)) +
//            double((y - right_dn_py) * (y - right_dn_py)));
//        int thing = con[ix][right_dn_py];
//        if (thing == 0) {
//            continue;
//        }
//        else {
//            tmp.push_back(make_pair(dis, thing));
//        }
//        con[ix][right_dn_py] = 0;
//    }
//
//    //下横行
//    for (int iy = right_dn_py; iy > left_up_py; iy--) {
//        double dis = sqrt(double((x - right_dn_px) * (x - right_dn_px)) +
//            double((y - iy) * (y - iy)));
//        int thing = con[right_dn_px][iy];
//        if (thing == 0) {
//            continue;
//        }
//        else {
//            tmp.push_back(make_pair(dis, thing));
//        }
//        con[right_dn_px][iy] = 0;
//    }
//    
//    //左竖行
//    for (int ix = right_dn_px; ix > left_up_px; ix--) {
//        double dis = sqrt(double((x - ix) * (x - ix)) +
//            double((y - left_up_py) * (y - left_up_py)));
//        int thing = con[ix][left_up_py];
//        if (thing == 0) {
//            continue;
//        }
//        else {
//            tmp.push_back(make_pair(dis, thing));
//        }
//        con[ix][left_up_py] = 0;
//    }
//
//    knife = recal_knife(tmp, knife);
//    tmp.clear();
//    if (left_up_px == 0 && left_up_py == 0 && right_dn_px == con.size() - 1 && right_dn_py == con.size()-1) {
//        return knife;
//    }
//    else {
//        return cal_length(con, x, y, depth + 1, knife);
//    }
//        
//}
//int main() {
//    set<int>muy;
//    muy.
//    int n = 0;
//    cin >> n;
//    cin.ignore();
//    while (n--) {
//        vector<vector<int>> con;
//        int size = 0;
//        int knife = 0;
//        cin >> size >> knife;
//        for (int i = 0; i < size; i++) {
//            vector<int> in_con;
//            int value = -1;
//            for (int j = 0; j < size; j++) {
//                cin >> value;
//                in_con.push_back(value);
//            }
//            con.push_back(in_con);
//        }
//        int x = 0, y = 0;
//        cin >> x >> y;
//        cin.ignore();
//        
//        cout<<"output begin()."<<endl;
//        for(auto item:con){
//            for(auto in_item:item){
//                cout<<in_item<<"-";
//            }
//            cout<<endl;
//        }
//        
//        cout << cal_length(con, x, y, 1, knife) << endl;;
//    }
//
//    return 0;
//}
