/*
 *功能：测试C++11中的thread的join, detach函数
 *作者：xyk
 *日期：2020.5.22
 */

 #include <iostream>
 #include <thread>

 using namespace std;

void func1()
{
    for(int i = 0; i<100; ++i)
    {
        cout<<"1";
    }
    cout<<"func1 end.";
}

void func2()
{
    for(int i = 0; i<100; ++i)
    {
        cout<<"2";
    }
    cout<<"func2 end.";
}

 int main()
 {
     
     thread do1(func1);
     do1.detach();
     thread do2(func2);
     do2.detach();
    for(int i = 0; i<100; ++i)
    {
        cout<<"m";
    }
    cout<<"main end.";
    
     return 0;
 }