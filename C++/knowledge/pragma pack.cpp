// pragma pack.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#pragma pack(show)
int main()
{
    std::cout << "Hello World!\n";
}

//可以看到下面编译的时候产生了一个warning显示对齐值是8
//https://www.jianshu.com/p/d994731f658d