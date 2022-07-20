#include "a.hpp"

int csum(int a, int b)
{
    return a+b+asum(a, b);
}
