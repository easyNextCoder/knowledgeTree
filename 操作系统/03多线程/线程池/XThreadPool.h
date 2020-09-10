#ifndef XTHREAD_POOL_H
#define XTHREAD_POOL_H

#include <vector>
#include <queue>
#include <memory>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <future>
#include <functional>
#include <stdexcept>

using namespace std;

class XThreadPool
{
public:
    XThreadPool(size_t size);
    template<class F, class... Args>
    auto enqueue(F&& f, Args&&... args) ->future<typename std::result_of<F(Args...)>::type>;
    ~XThreadPool();
private:
    vector<thread> workers;
    queue<function<void()>> tasks;
    mutex queue_mutex;
    condition_variable condition;
    bool stop;
};

XThreadPool::XThreadPool(size_t size):stop(false)
{
    for(int i = 0; i<size; ++i)
    {//这里开始启动，size个线程，一直取任务队列中得任务来运行
        workers.emplace_back([this,i]{//这一行不能使用push_back()，肯定底层调用的构造函数和emplace_back不同
            cout<<"start number: "<<i<<" thread."<<endl;
            for(;;)
            {
                std::function<void()> task;
                //使这个任务能够自动析构
                {
                    //由于任务进入队列和出队列要保持同步，所以使用condition_variable保持同步
                    unique_lock<mutex> lock(this->queue_mutex);
                    //如果队列没有停，或者队列是
                    this->condition.wait(lock, [this]{return this->stop || !this->tasks.empty();});
                    if(this->stop && this->tasks.empty())
                        return ;
                    task = move(this->tasks.front());
                    this->tasks.pop();
                }
                task();
            }
        });
    }
}

template<class F, class... Args>
auto XThreadPool::enqueue(F&&f, Args&&... args)->future<typename result_of<F(Args...)>::type>
{
    using return_type = typename result_of<F(Args...)>::type;
    //使用bind将任务绑定
    auto task = make_shared<packaged_task<return_type()>> (bind(forward<F>(f), forward<Args>(args)...));
    future<return_type> res = task->get_future();

    {
        //将任务加入队列和取出队列，需要保持同步，同一时刻只能发生二者之中的一件
        unique_lock<mutex> lock(queue_mutex);

        if(stop)
            throw runtime_error("enqueue on stopped ThreadPool");

        tasks.emplace([task](){(*task)();});
    }
    condition.notify_one();
    return res;
}
inline XThreadPool::~XThreadPool()
{
    {
        std::unique_lock<std::mutex> lock(queue_mutex);
        stop = true;
    }
    condition.notify_all();
    for(std::thread &worker: workers)
        worker.join();
}
#endif