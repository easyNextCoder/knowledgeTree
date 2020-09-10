#include <stdio.h>
#include <iostream>

using namespace std;

int main()
{
    int a[3] = {0, 1, 2}, *p = a;
    cout<<(int)a<<endl;
    cout<<(int)&a<<endl;    
    cout<<typeid(p).name()<<endl;
    cout<<typeid(a).name()<<endl;
    cout<<typeid(&a).name()<<endl;
    return 0;
}