/*
    date:2020.08.25
    for:测试类内和函数内的const变量是否必须要初始化
    ans:两者都必须要在定义的时候初始化
 */

#include <iostream>

using namespace std;

class A{
public:
    A(){};
private:
    const int b = 0;
};


int main()
{
    const int a = 0;

    return 0;
}