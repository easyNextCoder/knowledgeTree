#include <iostream>
#include <future>
#include <chrono>
 
void Thread_Fun1(std::promise<int> &p)
{
	//为了突出效果，可以使线程休眠5s
	std::this_thread::sleep_for(std::chrono::seconds(5));
 
	int iVal = 233;
	std::cout << "send in data:" << iVal << std::endl;
 
	//传入数据iVal
	p.set_value(iVal);
}
 
void Thread_Fun2(std::future<int> &f)
{
	//阻塞函数，直到收到相关联的std::promise对象传入的数据
	auto iVal = f.get();		//iVal = 233
 
	std::cout << "get data:" << iVal << std::endl;
}
 
int main()
{
	//声明一个std::promise对象pr1，其保存的值类型为int
	std::promise<int> pr1;
	//声明一个std::future对象fu1，并通过std::promise的get_future()函数与pr1绑定
	std::future<int> fu1 = pr1.get_future();
 
	//创建一个线程t1，将函数Thread_Fun1及对象pr1放在线程里面执行
	std::thread t1(Thread_Fun1, std::ref(pr1));
	//创建一个线程t2，将函数Thread_Fun2及对象fu1放在线程里面执行
	std::thread t2(Thread_Fun2, std::ref(fu1));
 
	//阻塞至线程结束
	t1.join();
	t2.join();
 
	return 1;
}