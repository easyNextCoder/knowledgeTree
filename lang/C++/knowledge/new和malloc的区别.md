## new和malloc的几点区别

### 申请的内存所在位置

new是在自由存储区上为对象动态分配内存空间，而malloc函数从堆上动态分配内存。

自由存储区是C++基于new操作符的一个抽象概念，凡是通过new操作符进行内存申请，该内存即为自由存储区。

#include <iostream>

using namespace std;

class testn{
public:
    string b;

};

int main()
{
    char a[100];
    a[99] = '\0';
    {
        testn * obj1 = new(a) testn();
        obj1->b = "yourname";
        cout<<obj1->b<<endl;    
    }
    cout<<a<<endl;
    return 0;
}

### 返回类型安全

new操作符内存分配成功时，返回的是对象类型的指针，类型严格与对象匹配，无需进行类型转换，所以new是符合类型安全的操作符。

### 内存分配失败时的返回值

new内存分配失败时，会抛出bad_alloc异常，它不会返回NULL；

malloc分配内存失败时，返回NULL

### malloc需要指定申请内存的大小，而new不需要

### 是否调用构造函数/析构函数

new操作符来分配对象内存时，会经历三个步骤

* 调用operator new函数，分配一块足够大，原始的，未命名的内存空间以便存储特定类型的对象

* 运行构造函数

* 返回一个指向该对象的指针

delete操作符释放对象内存时会经历两个步骤

* 调用对象的析构函数

* 用operator delete（或operator delete[]）函数释放内存空间

### operator new/ operator delete是可以被重载的

### 是否可以直观的重新分配内存

malloc分配内存后，如果使用过程中发现内存不足，可以使用realloc函数进行内存重新分配实现内存的扩充。realloc先判断当前的指针所指向的内存是否有连续的空间，如果有就原地扩大，如果没有就重新申请，之后拷贝过去。

new没有这样直观的配套设施来扩充内存

### 客户处理内存分配不足

new_handler set_new_handler(new_handler p) throw();

客户可以通过set_new_handler来设置自己处理内存分配不足时，应该怎么办



