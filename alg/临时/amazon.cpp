#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

const int offset = 1000001;

int main()
{
    int cnt = 0;
    cin>>cnt;
    vector<int> vec(offset, 0);
    for(int i = 0; i<cnt; ++i)
    {
        int val;
        cin>>val;
        vec[val]++;
    }

    //vector<int> dp(offset, 0);
    vector<vector<int>> dp(offset, vector<int>(2, 0));
    
    for(int i = 1; i<offset; ++i)
    {
        dp[i][0] = max(dp[i-1][1], dp[i-1][0]);
        dp[i][1] = max(dp[i-1][0]+i*vec[i], dp[i-1][1]);   
    }

    cout<<max(dp[offset-1][0], dp[offset-1][0]);
    return 0;
}