#include <iostream>

using namespace std;

int gcd(int a, int b)
{
    if(b == 0)
        return a;
        else
        {
            if(b == 1)
                cout<<"互质"<<endl;
            gcd(b, a%b);
        }
        
}

int main()
{
    cout<<gcd(31, 27)<<endl;
    return 0;
}