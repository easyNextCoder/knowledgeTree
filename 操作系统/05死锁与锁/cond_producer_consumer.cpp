//my_conditional_variable
#include <iostream>
#include <mutex>
#include <condition_variable>
#include <vector>
#include <random>
#include <thread>

using namespace std;

mutex m;
condition_variable cv;
const int MAX_SIZE = 10;
int container[MAX_SIZE];
int count = 0;
int thing = 0;

bool full = false;
bool empty = true;
default_random_engine e;
void producer()
{
    for(;;)
    {
        
        unique_lock<mutex> lk(m);
        cv.wait(lk, [](){return full == false;});//这句话的意思就是，如果不为满那么就开始生产
        //wait原地阻塞当前线程，并解锁lock，指导等到notify信号
        
        int produce_size = e()%MAX_SIZE+1;
        cout<<"producing+++++++++++++++++++++++++++++++++++++++++++"<<endl;
        std::this_thread::sleep_for(std::chrono::seconds(1));
        while(produce_size--)
        {
            container[count++] = thing++;
            if(count == MAX_SIZE)
            {
                full = true;
                break;
            }
        }
        empty = false;
        cout<<"produce done+++++++++++++++++++++++++++++++++++++++++"<<endl;
        //cout<<"call consume:"<<endl;
        //cv.notify_one();
        //cppreference
        //If any threads are waiting on *this, 
        //calling notify_one unblocks one of the waiting threads.
    }
}

void consumer()
{
    
    for(;;)
    {
        
        unique_lock<mutex> lk(m);
        cv.wait(lk, [](){return empty == false;});//这句话的意思是：如果不为空就开始消费
        
        int consume_size = e()%MAX_SIZE+1;
        std::this_thread::sleep_for(std::chrono::seconds(1));
        cout<<"consuming---------------------------------------------"<<endl;
        while(consume_size--)
        {
            if(count >= 0)
            {
                container[--count];
                cout<<container[count]<<" ";
            }else{
                empty = true;
                break;
            }
        }
        full = false;
        cout<<endl;
        cout<<"consuming done----------------------------------------"<<endl;
        //cv.notify_one();
    }
}

int main()
{
    
    thread t1(producer);
    thread t2(consumer);
    t1.join();
    t2.join();
    return 0;
}
