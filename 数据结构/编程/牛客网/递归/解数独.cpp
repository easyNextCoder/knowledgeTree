#include <iostream>

using namespace std;

const int N = 9;

int f[N][N];

pair<int, int> next(int i, int j)
{
    if(j<N-1)
    {
        j += 1;
    }else{
        i = i+1;
        j = 0;
    }
    return {i,j};
}

int dfs(int i, int j)
{
    //基本上所有的递归的失误的地方都是在处理终止条件的时候
    if(i == 9 && j == 0)return 1;
    
    if( f[i][j] != 0 )
    {
        pair<int, int> n = next(i, j);
        int rval = dfs(n.first, n.second);
        if(rval)return 1;
        else return 0;
    }else{
        //开始往0的地方填数
        int r[N] = {0,0,0,0,0,0,0,0,0};
        int m = (i/3)*3;
        int n = (j/3)*3;
        
        for(int p = 0; p<3; p++)
        {
            for(int q = 0; q<3; q++)
            {
                r[f[m+p][n+q]-1] = 1;
            }
        }
        for(int k = 0; k<N; k++)
        {
            r[f[i][k]-1] = 1;
            r[f[k][j]-1] = 1;
        }
        for(int l = N-1; l>=0; l--)
        {
            if(!r[l])
            {
                f[i][j] = l+1; 
                pair<int, int> n = next(i, j);
                
                int rval = dfs(n.first, n.second);
                if(rval)return rval;
                f[i][j] = 0;
            }
        }
        return 0;
    }
}

int main()
{
    for(int i = 0; i<N; ++i)
        for(int j = 0; j<N; ++j)
            cin>>f[i][j];
    dfs(0, 0);
    for(int i = 0; i<N; ++i)
    {
         for(int j = 0; j<N; ++j)
         {
             if(j == N-1)
                 cout<<f[i][j];
             else
                 cout<<f[i][j]<<" ";
         }
        cout<<endl;
    }
       
   
    return 0;
}

