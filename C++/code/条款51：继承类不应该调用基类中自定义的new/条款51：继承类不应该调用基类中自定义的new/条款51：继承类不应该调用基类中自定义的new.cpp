// 条款51：继承类不应该调用基类中自定义的new.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>

using namespace std;

class Base {
public:
    static void* operator new(std::size_t size) throw(bad_alloc);
    static void* operator new[](std::size_t size) throw(bad_alloc);
    int get_base_int() { return base_int; }
private:
    int base_int;
};

static void* operator new[](std::size_t size) throw(bad_alloc)
{
    if (size != 0) {
        cout << "this is new[]ing derived class." << endl;
        return ::operator new(size);
    }
    else {
        cout << "this is new[]ing base class." << endl;
        return malloc(size*sizeof(Base));
    }

}


void* Base::operator new(size_t size) throw(bad_alloc)
{
    if (size != sizeof(Base)) {
        cout << "this is newing derived class." << endl;
        return ::operator new(size);
    }
    else {
        cout << "this is newing base class." << endl;
        return malloc(size);
    }

}




class Derived :public Base {
public:
    int get_derived_int() { return derived_int; }

private:
    
    static int derived_int;
    const int dervied_const_int = 0;
    const int* derived_const_pint = 0;
};
int Derived::derived_int = -1;

int main()
{

    Base* pb = new Base();
    cout << pb->get_base_int() << endl;
    Derived* pd = new Derived();
    cout << pd->get_base_int() << endl;
    cout<<pd->get_derived_int() << endl;
    std::cout << "Hello World!\n";
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单
