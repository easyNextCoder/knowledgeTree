/*
 * 作用：测试C++与C中const变量的默认可链接属性不同
 * 作者：xyk
 * 日期：2020.5.23
 */
#include <stdio.h>
#include "head.h"

extern int sum(int a, int b);
const int varb = 999;
static int varc = 1000;
int main()
{
    int result = sum(10, 10);
    printf("the result is:%d\n", result);
    printf("value vara is:%d\n", vara);

}