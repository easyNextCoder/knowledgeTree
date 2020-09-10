#include <iostream>

using namespace std;

int main()
{
    int a[5] = {1,2,3,4,5};
    
    int *p = (int*)(&a+1);
    
    cout<<*(p-1)<<endl;
    return 0;
}