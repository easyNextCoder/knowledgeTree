// ���Ժ��麯������Ĵ�С.cpp : ���ļ����� "main" ����������ִ�н��ڴ˴���ʼ��������
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
    const double PI = 3.14159;//static �� const �������ã�Ϊʲô?

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

// ���г���: Ctrl + F5 ����� >����ʼִ��(������)���˵�
// ���Գ���: F5 ����� >����ʼ���ԡ��˵�

// ����ʹ�ü���: 
//   1. ʹ�ý��������Դ�������������/�����ļ�
//   2. ʹ���Ŷ���Դ�������������ӵ�Դ�������
//   3. ʹ��������ڲ鿴���������������Ϣ
//   4. ʹ�ô����б��ڲ鿴����
//   5. ת������Ŀ��>���������Դ����µĴ����ļ�����ת������Ŀ��>�����������Խ����д����ļ���ӵ���Ŀ
//   6. ��������Ҫ�ٴδ򿪴���Ŀ����ת�����ļ���>���򿪡�>����Ŀ����ѡ�� .sln �ļ�

