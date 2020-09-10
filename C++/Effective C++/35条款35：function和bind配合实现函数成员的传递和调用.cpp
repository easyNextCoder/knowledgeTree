// 条款35：function和bind配合实现函数成员的传递和调用.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <functional>

using namespace std;

class GameCharacter;
int defaultHealthCalc(const GameCharacter& gc) 
{
    cout << "I am defaultHealthCalc." << endl;
    return 1;
}

typedef function<int(const GameCharacter&)>HealthCalFunc;
class GameCharacter {
    
public:
    explicit GameCharacter(HealthCalFunc hcf = defaultHealthCalc)
        :healthFunc(hcf)
    {

    }

    int healthValue() const
    {
        return healthFunc(*this);
    }
private:
    HealthCalFunc healthFunc;
};

class GameLevel {
public:
    int health(const GameCharacter&) {
        cout << "I am GameLevel health." << endl;
        return 2;
    };
};

class evilGuy:public GameCharacter
{
public:
    evilGuy(HealthCalFunc hcf = defaultHealthCalc) :GameCharacter(hcf)//构造函数一定不能少，而且如果前面不设置public关键字，则对象无法构造
    {

    }
};

//尝试在对象类中返回成员函数指针仍然错误：&绑定成员函数表达式上的非法操作
/*

class A {
public:

    int get() {};
    function<int()> return_getf() {
        return bind(&(A::get), *this);
    }

};

*/



int main()
{
    GameLevel evilLevel;
    evilGuy eg(bind(&GameLevel::health, evilLevel, placeholders::_1));
    cout << "eg.healthValue is:" << eg.healthValue() << endl;

    return 0;
}
