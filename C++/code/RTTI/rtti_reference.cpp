#include <iostream>

using namespace std;

class CB{
public:
    int getInt(){return 1;}
};

class CI:public CB{

};
int main()
{
    CB* p = new CI();
    CB &t = *p;

    cout<<typeid(CB*).name()<<endl;
    cout<<typeid(*p).name()<<endl;
    cout<<typeid(CB&).name()<<endl;
    cout<<typeid(t).name()<<endl;
    cout<<typeid(CI).name()<<endl;

    CI* derived = new CI();
    CB* base = new CB();
    //base = static_cast<CB*>(derived);
    //derived = static_cast<CI*>(base);
    derived = dynamic_cast<CI*>(base);

    return 0;
}