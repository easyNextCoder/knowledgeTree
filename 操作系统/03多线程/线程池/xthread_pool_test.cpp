#include <iostream>
#include "XThreadPool.h"

using namespace std;

int main()
{
    XThreadPool xw(5);

    for(int i = 0; i<10; ++i)
    {
        xw.enqueue([i]{
            for(int j = 0; j<10000; ++j)
                std::cout<<i;
        });
    }

    return 0;
}