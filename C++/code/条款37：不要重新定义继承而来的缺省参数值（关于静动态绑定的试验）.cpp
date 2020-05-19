#include<bits/stdc++.h>
using namespace std;
class Base
{
public:
    virtual void f()
    {
        cout << "Base::f" << endl;
    }
    virtual void g()
    {
        cout << "Base::g" << endl;
    }
    virtual void h()
    {
        cout << "Base::h" << endl;
    }

};
class Base_son:public Base
{
public:
    virtual void f()
    {
        cout << "Base_son::f" << endl;
    }
    virtual void g1()
    {
        cout << "Base_son::g1" << endl;
    }
    virtual void h1()
    {
        cout << "Base_son::h1" << endl;
    }

};


typedef void(*Fun)(void);
Base_son d;

Fun pFun = NULL;

int main()
{
    cout << "虚函数表地址：" << (int*)(&d) << endl;
    cout << "虚函数表 — 第一个函数地址：" << (int*)*(int*)(&d) << endl;

    //通过虚函数表调用虚函数
    pFun = (Fun)*((int*)*(int*)(&d));   // Base_son::f()
    pFun();
    pFun =(Fun)*((int*)*(int*)(&d)+1);  // Base::g()
    pFun();
    pFun =(Fun)*((int*)*(int*)(&d)+2);  // Base::h()
    pFun();

    pFun =(Fun)*((int*)*(int*)(&d)+3);  // Base_son::g1()
    pFun();
    pFun =(Fun)*((int*)*(int*)(&d)+4);  // Base_son::h1()
    pFun();

	Base a;
	a.f(); 
	Base *b=new Base_son(); b->f();

    return 0;
}
