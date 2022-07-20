//joinable

#include <iostream>
#include <thread>

using namespace std;

int main()
{
    
        thread t1 = thread([]{while(1)cout<<"hello1"<<endl;});
        //thread t2 = thread([]{cout<<"hello2"<<endl;});
        //t1.join();
        this_thread::sleep_for(chrono::seconds(3));
        t1.~thread();
        cout<<"main end"<<endl;
        this_thread::sleep_for(chrono::seconds(3));
       //cout<<"output t1:"<<t1.joinable()<<endl;
        
    
    return 0;
}

