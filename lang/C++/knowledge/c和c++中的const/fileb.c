#include "head.h"
//extern int vara;
extern const int varb;
//extern static int varc;
//这个varc变量无法声明，因为他是一个static的
int sum(int a, int b)
{
    vara = 99;
    return a+b+vara;
}