#include <stdio.h>

int main()
{
    int a[10];
    int (*b)[10];
    int *c = a;
    printf("sizeof a:%d.\n", sizeof(a));
    printf("sizeof b:%d.\n", sizeof(b));
    printf("sizeof *b:%d.\n", sizeof(*b));
    printf("sizeof c:%d.\n", sizeof(c));
    printf("sizeof *c:%d.\n", sizeof(*c));
    //https://blog.csdn.net/findgeneralgirl/article/details/78501734
    //以下三个值是相同的
    printf("addr of a:%d\n", a);
    printf("addr of &a:%d\n", &a);
    printf("addr of &a[0]:%d\n", &a[0]);
    
    return 0;
}