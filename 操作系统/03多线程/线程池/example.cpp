#include <iostream>
#include <vector>
#include <chrono>

#include "ThreadPool.h"

using namespace std;

int main()
{
    
    ThreadPool pool(4);
    std::vector< std::future<int> > results;

    for(int i = 0; i < 100; ++i) {
        results.emplace_back(
            pool.enqueue([i] {
                std::cout << "hello " << i << std::endl;
                std::this_thread::sleep_for(std::chrono::seconds(1));
                std::cout << "world " << i << std::endl;
                return i*i;
            })
        );
        cout<<"now task size is:"<<pool.tack_size()<<endl;
        cout<<"working thread size is:"<<pool.working_thread_size()<<endl;
    }

    for(auto && result: results)
        std::cout << result.get() << ' ';
    std::cout << std::endl;
    
    return 0;
}