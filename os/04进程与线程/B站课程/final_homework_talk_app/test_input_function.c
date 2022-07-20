#include <stdio.h>
#include <stdlib.h>

int main(){

    char * s = (char*)malloc(sizeof(char)*100);
    while( gets(s) != EOF){

        printf("%s\n",s);

    }
    
    return 0;
}