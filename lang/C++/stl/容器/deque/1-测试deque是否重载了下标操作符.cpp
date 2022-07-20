#include <iostream>
#include <deque>

using namespace std;

int main()
{

    deque<int> dq = {1,2,3,4,5,6};
    cout<<dq[0]<<endl;
    cout<<dq[dq.size()-1]<<endl;

    return 0;
}