#include <iostream>

using namespace std;

class testn{
public:
    string b;

};

int main()
{
    char a[100];
    a[99] = '\0';
    {
        testn * obj1 = new(a) testn();
        obj1->b = "yourname";
        cout<<obj1->b<<endl;    
    }
    cout<<a<<endl;
    return 0;
}