#include <iostream>
#include <vector>
#include <stdio.h>
#include <cstring>

using namespace std;

const int N = 1001;

int f[N][N];


vector<vector<int>> input;
int getMax(int n, int V)
{
    if(f[n][V] != -1)return f[n][V];
    if(n == input.size() || V<0)
        return 0;
    int maxOut = 0;
   
    //选择这个加入背包
    int rval1 = 0;
    if(V-input[n][0] >= 0)
        rval1 = getMax(n+1, V-input[n][0])+input[n][1];
        //跟动态规划里面的看着一样，不过动态规划是取上面一次
        //的结果，而递归是再运算一下得到结果，不过如果把结果保存
        //下来实际效果差不多，几倍时间的差距
    //这个不选择加入背包
    int rval2 = getMax(n+1, V);

    f[n][V] = max(rval1, rval2);
    return max(rval1, rval2);
}

int main()
{
    int n, m;
    cin>>n>>m;
    memset(f, -1, sizeof(f));
    for(int i = 0; i<n; ++i)
    {
        int tmpv, tmpm;
        cin>>tmpv>>tmpm;
        input.push_back({tmpv, tmpm});
    }
    cout<<getMax(0, m)<<endl;
    return 0;
}

