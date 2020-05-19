// 条款52：测试自己的类中的new覆盖global中的new.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <string>
#include <iterator>
#include <algorithm>
#include <vector>

using namespace std;


class B {
public:
    virtual void f() const;
};

class D :public B {
public:
    virtual void f();
};


class base {
public:
    /*
        //自己的基类中写一个base类new，会将global中的
        //void* operator new(std::size_t)throw(std::bad_alloc);//normal版本
        //void* operator new(std::size_t, void*)throw();       //placement new
        //void* operator new(std::size_t, const std::nothrow_t&) throw()//nothrow new
        这三版都会被覆盖
    */

    void* operator new(size_t size) throw(bad_alloc);
};

void* base::operator new(size_t size) throw(bad_alloc) {
    ;
    return 0;
}


int main()
{

    /*
    void* vp = NULL;
    base* p = new(vp) base;//如果自己重定义了New,这里就会报错
    cerr << "test error." << endl;
    clog << "test log" << endl;
    std::cout << "Hello World!\n";
    */
    
    string s("myname");
    ostream_iterator<string>out(cout);
    *out++ = s;
    *out++ = "\n";

    
    istream_iterator<string>myin(cin);
    /*
    while (!myin->empty()) {
        cout << *(++myin) << endl;
    }
    */

    for_each(myin, ++myin, [](string s) {cout << "your name."<<s << endl; });//?
   
    //测试传入for_each中的表达式是否可以是引用类型
    vector<int> vec(10, 1);
    for_each(vec.begin(), vec.end(), [](int& a) {cout << a << endl; a = 2; });
    for (auto item : vec) {
        cout << item << endl;
    }


    return 0;
}

