/*  
    data:2020.6.27
    仔细看了一遍现在C++中mutex已经集成了lock所述的功能
 */


#include <iostream>
#include <string>
#include <vector>
#include <mutex>

using namespace std;

class automicClass{
pubic:
    automicClass()
    {
        m.lock();
    }

    ~automicClass()
    {
        m.unlock();
    }
private:
    mutex m;
};


int main()
{

    return 0;
}