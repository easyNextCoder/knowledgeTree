#include <stdio.h>
#include "a.h"

const int varb = 10;
/*
    在C语言中const的默认属性是外部链接属性的
    
 */

int main(){

    printf("hello.\n");
    printf("the varb is:%d\n", getVarb());
    printf("the sum2a is:%d\n", sum2a());
    printf("the sum is:%d\n", sum(a, 2));
    return 0;

}

