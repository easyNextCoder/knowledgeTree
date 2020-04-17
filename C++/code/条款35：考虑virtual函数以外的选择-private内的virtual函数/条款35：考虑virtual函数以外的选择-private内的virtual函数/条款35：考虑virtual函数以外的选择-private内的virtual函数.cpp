// 条款35：考虑virtual函数以外的选择-private内的virtual函数.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <string>

using namespace std;

class GameCharacter {
public:
    int healthValue() const
    {
        int retVal = doHealthValue();
        return retVal;
    }
protected:
    int getsomething() {
        return 100;
    }
private:
    virtual int doHealthValue() const
    {
        return 10;
    }
};

class spiderMan :private GameCharacter {
public:
    int rval = GameCharacter::getsomething();
private:
    virtual int doHealthValue() const {
        return 100;
    }
};

int main()
{
    spiderMan awey;
    cout << awey.healthValue() << endl;;
    std::cout << "Hello World!\n";
}
