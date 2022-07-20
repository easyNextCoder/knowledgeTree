#include <iostream>


using namespace std;

const int N = 310;
int dp[N][N];
int stone[N];

int getMin(int first, int last)
{
    if(last -first+1 <=  2)
    {
        dp[first][last] = stone[first]+stone[last];
        return dp[first][last];
    }
    for(int i = first; i<= last; i++)
    {
        for(int len = 1; len+first<=last; len++)
        {
            int j = first + len;
            for(int k = i; k<=j; k++)
            {
                dp[i][j] = max(dp[i][j], dp[i][k]+dp[j][k]);
            }
        }
    }
    return dp[first][last];
}

int main()
{
    int n;
    cin>>n;
    for(int i = 1; i<=n; i++)
        cin>>stone[i];
    getMin(1, n);
    
    
    return 0;
}
https://www.acwing.com/problem/content/284/