#include <stdio.h>

int main(){

    char *s = malloc(sizeof(char)*100);
    memcpy(s, "1234567890", 10);
    printf("%s  %x\n", s, s);

    s = s+4;

    printf("%s  %x\n", s, s);
    

    return 0;
}