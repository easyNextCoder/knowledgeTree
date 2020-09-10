#include <iostream>

using namespace std;

typedef union {
    int m;
    char a[10];
    double f;
} DATA;

typedef union{
    int m;
    char a[9];
    int n;
}DATA1;

struct {
    int m;
    DATA k; 
    double p;
}real;


int main()
{
    cout<<sizeof(DATA)<<endl;
    cout<<sizeof(DATA1)<<endl;
    cout<<sizeof(real)<<endl;
    return 0;
}