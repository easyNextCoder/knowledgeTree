
#include <iostream>

using namespace std;

class AClass {
public:
    AClass() = default;
    AClass(AClass& input)
    {
        cout << "\tinvoking copy constructor." << endl;
    }
    AClass operator=(AClass& input)
    {
        cout << "\tinvoking = constructor." << endl;
        return *this;
    }
    /*
    AClass& operator=(AClass& input)
    {
        cout << "\tinvoking &= constructor." << endl;
        return *this;
    }
    */
    
private:

};

int an_main()
{
    cout << "这里调用的是default constructor." << endl;
    AClass a;
    
    cout << "这里为什么只调用copy constructor." << endl;
    AClass b = a;
    
    cout << "因为是使用引用赋值，所以只调用= constructor:" << endl;
    AClass &c = a;
    c = b;

    return 0;
}

