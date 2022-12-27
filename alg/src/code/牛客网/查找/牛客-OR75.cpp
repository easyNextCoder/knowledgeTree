#include <iostream>

using namespace std;

const int N = 1002;
int map[N][N];
int dp[N][N];
//dp[i][j] 表示从坐标(0,0)走到坐标(i,j)所需要的最小行动力

int dp1[N][N];
//当前运行值
int main()
{
    int m, n;
    cin>>m>>n;
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
             cin>>map[i][j];   
        }
    }
    
    if(map[0][0]>=0)
    {
        dp[0][0] = 0;
        dp1[0][0] = map[0][0];
    }
    else
    {
        dp1[0][0] = 1;
        dp[0][0] = 1-map[0][0];
    }
        
    for(int i = 1; i<n; i++)
    {
        if(dp1[0][i-1] + map[0][i]>0)
        {
            dp[0][i] = dp[0][i-1];
            dp1[0][i] = dp1[0][i-1]+map[0][i];
        }
        else{
            dp[0][i] = dp[0][i-1] - (dp1[0][i-1]+map[0][i])+1;
            dp1[0][i] = 1;
        }
    }
    
    for(int j = 1; j<m; j++)
    {
        if(dp1[j-1][0] + map[j][0]>0)
        {
            dp[j][0] = dp[j-1][0];
            dp1[j][0] = dp1[j-1][0]+map[j][0];
        }
        else{
            dp[j][0] = dp[j-1][0]-(dp1[j-1][0]+map[j][0])+1;
            dp1[j][0] = 1;
        }
    }
    
    for(int i = 1; i<m; i++)
    {
        for(int j = 1; j<n; j++)
        {
            int dpmin1 = 0;
            if(dp1[i-1][j] + map[i][j] > 0)
            {
                dpmin1 = dp[i-1][j];
            }else{
                dpmin1 = dp[i-1][j]-(dp1[i-1][j]+map[i][j])+1;
            }

            int dpmin2 = 0;
            if(dp1[i][j-1] + map[i][j] > 0)
            {
                dpmin2 = dp[i][j-1];
            }else{
                dpmin2 = dp[i][j-1]-(dp1[i][j-1]+map[i][j])+1;
            }
            
            int dpmin = min(dpmin1, dpmin2);
            int dp1min1 = 0;
            int dp1min2 = 0;
            if(dpmin == dpmin1)
            {
                if(dp1[i-1][j]+map[i][j] > 0)
                {
                    dp1min1 = dp1[i-1][j]+map[i][j];
                }else{
                    dp1min1 = 1;
                }   
                dp1[i][j] = dp1min1;
            }else{
                if(dp1[i][j-1]+map[i][j]>0)
                {
                    dp1min2 = dp1[i][j-1]+map[i][j];
                }else{
                    dp1min2 = 1;
                }
                dp1[i][j] = dp1min2; 
            }
            dp[i][j] = dpmin;
        }
    }
    
    cout<<endl;
    cout<<"the dp result is: "<<endl;
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
            cout<<dp[i][j]<<" ";
        }
        cout<<endl;
    }

    cout<<"dp1 result is:"<<endl;
    cout<<endl;
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
            cout<<dp1[i][j]<<" ";
        }
        cout<<endl;
    }
    
    cout<<dp[m-1][n-1]<<endl;
    return 0;
}