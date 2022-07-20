#include <iostream>
#include <string>

using namespace std;

class CBase
{
public:
    //纯虚函数也可以定义也可以不定义，但是继承类必须继承接口
    //virtual 可以继承接口也可以继承函数体
    //实体函数必须全部继承所有的东西
    virtual char* getName() = 0;
};

class CBint:public CBase
{
public:
    char* getName(){return "CBint";}
    int getInt(){return 1;}
};

class CBString:public CBase
{
public:
    char* getName(){return "CBString";}
    char* getString(){return "string";}
};

int main()
{
    CBase* B1 = (CBase*)new CBint();
    cout<<B1->getName()<<endl;

    CBint* B2 = static_cast<CBint*>(B1);
    if(B2)
    {
        cout<<B2->getInt()<<endl;
    }

    CBase* C1 = (CBase*)new CBString();
    cout<<C1->getName()<<endl;

    CBString * C2 = static_cast<CBString*> (C1);
    if(C2)
    {
        cout<<C2->getString()<<endl;
    }
    return 0;
}