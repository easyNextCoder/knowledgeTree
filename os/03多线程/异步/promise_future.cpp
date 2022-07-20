#include <iostream>
#include <future>
#include <thread>
#include <string>
#include <vector>
#include <functional>

using namespace std;
using std::string;

string scat(string s1, string s2)
{
    return s1+s2;
}

using FT = function<string (string, string)>;

void thread_fun0(promise<FT> & f)
{
    cout<<"I am thread_fun0, going to set_value."<<endl;
    f.set_value(bind(&scat, std::placeholders::_1, std::placeholders::_2));
}

/*
template<typename T, typename ...Args>
void get(T& a, Args&& ...args)
{
    forward<Args>(args)...
}
*/
template<typename T, typename ...Args>
void thread_fun1(future<FT> & f, Args&& ...args)
{
    cout<<"I am thread_fun1, going to get."<<endl;
    auto ff = f.get();
    auto result = ff(std::forward<Args>(args) ...);
    cout<<result<<endl;
}

template<typename T>
int complexFunc(bool a, int b, string c, T d)
{
    cout<<a<<" "<<b<<" "<<c<<" "<<d<<endl;
    return a;
}

int main()
{
    promise<FT> prf;
    future<FT> frf = prf.get_future();
    string s1 = "xu";
    string s2 = "yong";
    auto t1 = thread(thread_fun0, ref(prf));
    auto t2 = thread(thread_fun1<FT, string, string>, ref(frf), s1, s2);
    t1.join();
    t2.join();
    
    //测试给任务做一个wrapper
    //auto task1 = []{complexFunc(true, 10, "xuyongkang", 'x');};
    auto task2 = []{while(1){cout<<"task2"<<endl;};};
    //测试建立任务容器，任务容器中能放置各种任务
    vector<thread> tasks;
    //tasks.push_back(task1);
    tasks.emplace_back(task2);
    thread tmp([]{while(1){cout<<"tmp"<<endl;}});
    tasks[0] = tmp;//thread([]{while(1){cout<<"task3."<<endl;};});
    //(tasks.back())();
    //(tasks.front())();
    return 0;
}

