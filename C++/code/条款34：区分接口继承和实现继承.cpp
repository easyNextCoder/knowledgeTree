// 条款34：区分接口继承和实现继承.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
#include <iostream>

using namespace std;

class Airport {

};

class Airplane {
public:
    virtual void fly(const Airport& destination) = 0;
    virtual void specialFly() ;
    void coreFly();
};

void Airplane::fly(const Airport& destination)//pure virtual函数实现
{
    //缺省行为，将飞机飞至指定目的地
    cout << "this is Airplane default fly." << endl;
}
void Airplane::specialFly() {
    cout << "this is Airplane specialFly." << endl;
}
void Airplane::coreFly() {
    cout << "this is Airplane coreFly." << endl;
}


class ModelA :public Airplane {
public:
    virtual void fly(const Airport& destination) {
        Airplane::fly(destination);
    }
};

class ModelB :public Airplane {
public:
    virtual void fly(const Airport& destination) ;//强制继承纯虚的接口（可以实现继承类的默认行动），可以选择调用基类中的default行动
    virtual void specialFly();                    //指定接口继承缺省实现继承（即如果自己没有定义specialFly行为，base将提供一份缺省的）
    //void coreFly();                             //强制继承类全部继承了基类的接口和实现，一定不要在继承类中重新定义，因为他代表不变性和特殊性
};

void ModelB:: fly(const Airport& destination) {
    Airplane::fly(destination);
    cout << "this is ModelB fly." << endl;
}

void ModelB::specialFly() {
    cout << "this is ModelB specialFly." << endl;
}

/*
void ModelB::coreFly() {
    cout << "this is ModelB coreFly." << endl;
}
*/
int main()
{
    ModelA a_plane;
    ModelB b_plane;
    a_plane.fly(*(new Airport));
    a_plane.specialFly();

    b_plane.fly(*(new Airport));
    b_plane.specialFly();
    b_plane.coreFly();


    std::cout << "Hello World!\n";
}
