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
    int length;
    int width;
};

class Circle  {
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
    const double PI = 3.14159;//static 和 const 不能连用，为什么?

    int radius = 0;

	double volum = 0;
	
	char ID = 'c';
};

int main()
{
    BaseV item;
    
    Circle c;

    c.setRadius(10);
    cout << c.getArea() << endl;


    cout << sizeof(item) << endl;
    cout << sizeof(c) << endl;
    cout << sizeof(double) << endl;
    
    std::cout << "Hello World!\n";
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件

