
#include <iostream>

using namespace std;
unsigned int con[32];

int main()
{
    int a, b;
    cin>>a>>b;
    unsigned int one = 1;
    if(!( a>=1 && a<=1024 && b>=1 && b<=1024 ))
    {
        cout<<-1<<endl;
        return 0;
    }
    a--;
    b--;
    
        
    con[a/32] = con[a/32] | (one<<(31-(a%32)));
    
    if(( con[b/32] & (one<<(31-(b%32))) ) > 0)//优先级问题难倒门内汉
        cout<<1<<endl;
    else
        cout<<0<<endl;
    return 0;
}