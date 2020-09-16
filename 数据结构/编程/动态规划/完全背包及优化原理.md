> https://www.acwing.com/problem/content/3/

## 完全背包的朴素写法：

````
#include <iostream>

using namespace std;

const int N = 1010;
int n, m;
int v[N], w[N];
int f[N][N];

int main()
{
    cin>>n>>m;
    for(int i = 1; i<=n; i++)cin>>v[i]>>w[i];
    
    for(int i = 1; i<=n; i++)
    {
        for(int j = 0; j<=m; j++)
        {
            /*
            for(int k = 0; k*v[i]<=j; k++)
                f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i]);
            */
            f[i][j] = f[i-1][j];
            if(j>=v[i])f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i]);
        }
    }
    
    cout<<f[n][m]<<endl;
}

````

## 完全背包的优化写法及原理

````
#include <iostream>

using namespace std;

const int N = 1010;
int n, m;
int v[N], w[N];
int f[N];

int main()
{
    cin>>n>>m;
    for(int i = 1; i<=n; i++)cin>>v[i]>>w[i];
    
    for(int i = 1; i<=n; i++)
    {
        for(int j = 0; j<=m; j++)
        {
            /*
            for(int k = 0; k*v[i]<=j; k++)
                f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i]);
            */
            //f[i][j] = f[i-1][j];//在这里本该被替换为f[j] = f[j]可以省略不写，可以替换的原因是：由于是从小到达进行的运算在f[j] = f[j]等号右边的f[j]在这个等式中还没有被更新，也就是f[j] = f[i-1][j],
            if(j>=v[i])f[j] = max(f[j], f[j-v[i]]+w[i]);
            //也是由于是从小到达进行的计算，所以，f[j-v[i]]使用的是更新过的值，这正是我们想要的
        }
    }
    
    cout<<f[m]<<endl;
}

````