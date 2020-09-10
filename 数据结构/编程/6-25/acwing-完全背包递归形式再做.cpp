#include <iostream>

using namespace std;

const int N = 1010;
int n, m;
int v[N], w[N];
int f[N];

int getMax(int index, int leftV)
{
    if(index == n ||leftV == 0)return 0;
    int rval = 0;
    for(int i = 0; i*v[index]<=leftV; i++)
    {
        rval = max(rval, getMax(index+1, leftV-i*v[index])+i*w[index]);
    }
    return rval;
}

int main()
{
    cin>>n>>m;
    for(int i = 0; i<n; i++)cin>>v[i]>>w[i];
    cout<<getMax(0, m)<<endl;
    return 0;
    // for(int i = 1; i<=n; i++)
    // {
    //     for(int j = 0; j<=m; j++)
    //     {
    //         /*
    //         for(int k = 0; k*v[i]<=j; k++)
    //             f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i]);
    //         */
    //         //f[i][j] = f[i-1][j];
    //         if(j>=v[i])f[j] = max(f[j], f[j-v[i]]+w[i]);
    //     }
    // }
    
    // cout<<f[m]<<endl;

    
}