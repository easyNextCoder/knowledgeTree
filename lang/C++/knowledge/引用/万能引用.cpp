/*
 * date:2020.8.5
 * for:什么是万能引用？
 * 
 */
#include <iostream>
#include <string>

using namespace std;

//普通的万能引用用在函数中
template<typename T>
void fun(T&& var)
{
    cout<<var<<endl;
}

//类内的成员实现万有引用
template<typename T>
class uref{
public: 
    template<typename M>
    void uref_fun(M&& out)
    {
        cout<<out<<endl;
    } 
};

int main()
{
    //测试万能引用参数函数
    fun(100);
    int i = 200;
    fun(i);
    fun(string("hello world"));
    //测试万能引用在类的成员函数中
    uref<int> a;
    a.uref_fun(string("name"));
    return 0;
}