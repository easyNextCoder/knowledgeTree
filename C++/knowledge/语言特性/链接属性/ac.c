#include "a.h"

int sum(int a, int b)
{
    return a+b;
}

int sum2a(void)
{
    return a+a;
}

extern int varb;
int getVarb()
{
    return varb;
}

