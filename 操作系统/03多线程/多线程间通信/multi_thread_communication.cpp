#include <iostream>
#include <vector>
#include <thread>
#include <mutex>

using namespace std;

const int N = 100;
mutex m,n;

class common_data{

public:
    void get(){
        n.lock();
        m.lock();
        for(int i = 0; i<N; ++i)
        {
            cout<<value[i];
        }
        cout<<endl;
        n.unlock();
        m.unlock();
        //return value;
    };
    void set(int val){
        m.lock();
        for(int i = 0; i<N; ++i)
            value[i] = val; 
        m.unlock();
    };

private:
    int value[N];
};

int main()
{
    common_data data_a;
    
    auto fun1 =  [&data_a](){
            while(1)
            {
                
                data_a.set(1);
                
            }
        };
    auto fun2 = [&data_a](){
            while(1)
            {
                data_a.set(9);
                //data_a.get();
                //cout<<"thread2:"<<(2 == data_a.get())<<" ";
            }
        };
    
    auto fun3 = [&data_a]()
    {
        while(1)
        {
            data_a.get();
        }
    };

    auto fun4 = [&data_a]()
    {
        while(1)
        {
            data_a.get();
        }
    };

    thread a(fun1);
    thread b(fun2);
    thread c(fun3);
    thread d(fun4);

    a.join();
    b.join();
    c.join();
    d.join();
    
    return 0;
}