// 条款39：测试继承类是否可以访问私有继承的基类中私有的成员.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
using namespace std;

class A {
public:
    void get() {
        cout << "get()" << endl;
    }
private:
    void private_get() {
        cout << "private get()" << endl;
    }
};

class B : private A {
public:
    void bget() {
        A::get();
       
    }
private:
    void b_private_get() {
        A::get();//不能访问A中的私有函数
    }
};

int main()
{
    std::cout << "Hello World!\n";
}
