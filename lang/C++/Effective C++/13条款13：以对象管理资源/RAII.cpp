/*
    data:2020.6.27
    RAII-取得资源时便进行初始化-Resource Acquisition Is Initialization

 */

#include <iostream>
#include <memory>

using namespace std;

class Investment{
public:

    ~Investment()
    {
        cout<<"I am destorying."<<endl;
    }
private:
    int var;
};

Investment* createInvestment(){

    return new Investment();
}

void dele(Investment* a)
{
    cout<<"I am dele."<<endl;
}

void f()
{
    shared_ptr<Investment> pInv1(createInvestment(), dele);

    shared_ptr<Investment> pInv2(pInv1);

    pInv1 = pInv2;
}

int main()
{
    f();
    return 0;
}