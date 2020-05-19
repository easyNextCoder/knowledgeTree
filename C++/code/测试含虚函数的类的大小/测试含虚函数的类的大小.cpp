// 测试含虚函数的类的大小.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>

using namespace std;



class BaseV {
public:
    BaseV& setLength(int value)
    {
        length = value;
        return *this;
    }

    int getLength(int value)
    {
        return length;
    }

    BaseV setWidth(int value)
    {
        width = value;
        return *this;
    }

    int getWidth(int value)
    {
        return width;
    }

    virtual double getArea()
    {
        return length * width;
    }

private:
    char length;
    char width;
    double a;
};

class Rectangle :private BaseV {
public:
    double getArea() {
        return 0;
    }
private:
    int lwRatio = 0;
};

class Circle {
public:
    Circle& setRadius(double value)
    {
        radius = value;
        return *this;
    }

    double getRadius()
    {
        return radius;
    }

    virtual int getArea()
    {
        return radius * radius * PI;
    }

private:
    const int PI = 3.14159;//static 和 const 不能连用，为什么?

    int radius = 0;

    int volum = 0;

    char ID = 'c';

    char ID2;//从1-4个char最终的结果都是占4个字节
    char ID3;
};

int main()
{
    BaseV item;

    Circle c;

    Rectangle r;


    cout << sizeof(item) << endl;
    cout << sizeof(c) << endl;
    cout << sizeof(r) << endl;
    
    std::cout << "Hello World!\n";
}