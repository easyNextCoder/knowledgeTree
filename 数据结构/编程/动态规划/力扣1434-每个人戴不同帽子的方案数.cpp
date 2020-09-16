//这是使用回溯+记忆搜索的方法做的，如果不加上记忆搜索的话，严重超时！
#include <iostream>
#include <vector>
#include <cstring>

using namespace std;

class Solution {
public:
    long long int one = 1;
    
    vector<vector<int>> people2hats = vector<vector<int>>(41, vector<int>());
    int dp[1025][41];
    int finalState = 0;
    int mod = 1e9 + 7;
    int dfs(int u, int state)
    {
       
        if(state == finalState)
        {
            return 1;
        }
        if(u >= 41)
        {
            return 0;
        }
        //最后排查了好几个小时的超时问题就是这里
        //起初写成了if(dp[state][u]){return dp[][];}
        //这样需要计算许多等于0的情况，所以超时
        
        if(dp[state][u] != -1)
        {
            return dp[state][u];
        }

        int count =  dfs(u+1, state);
        auto it = people2hats[u].begin();
        while(it != people2hats[u].end())
        {
            if((state & one <<*it) == 0)
            {
                count = ( count + dfs(u+1, (state|(one<<*it))))%(mod);
            }  
            ++it;
        }
        
        dp[state][u] = count;
        return count;         
    }

    int numberWays(vector<vector<int>>& hats) {
        memset(dp, -1, sizeof(dp));
        int cnt = 0;
        for (int i = 0; i < hats.size(); ++i) {
			for (int j = 0; j < hats[i].size(); ++j) {
                people2hats[hats[i][j]].push_back(i);
            }
            finalState |= 1<<(cnt);
            cnt++;
        }
        cout<<finalState<<endl;
        //最终要有hats.size()个人获得帽子，才算是一种情况
        return dfs(0, 0);
    }
};

int main()
{
    Solution so;
    vector<vector<int>> vec = {{1,2,3},{2,3,5,6},{1,3,7,9},{1,8,9},{2,5,7}};
    so.numberWays(vec);
    return 0;
}

/*//以下是一个比较好的使用转移方程做的！
解题思路
状态压缩DP

定义状态dp[i][j] 为从1 - i将1 - i顶帽子，分给状态为j的方法。则分为两类：

不使用第i顶帽子，则为dp[i-1][j]
使用第i顶帽子，则枚举状态j中的人，并且这个人也喜欢这个帽子，有dp[i-1][j - (1 << k)], 第k个人待第i顶帽子的方法
最终输出dp[40][(1 << n) - 1]

时间复杂度40 \times 2 ^n \times n40×2 
n
 ×n
空间复杂度40 \times 2 ^ n40×2 
n
 

个人博客

代码

class Solution {
public:
    
    int dp[41][1 << 10] = {0};
    bool has[10][41] = {false};
    const int mod = 1e9+7;
    
    int numberWays(vector<vector<int>>& hats) {
        int n = hats.size();
        for(int i = 0; i < n; i ++){  // 预处理好第i个人是否喜欢第j顶帽子，用于后面转移判断使用
            for(auto j : hats[i]){
                has[i][j] = true;
            }
        }
        dp[0][0] = 1;
        for(int i = 1; i <= 40; i ++){
            for(int j = 0; j < 1 << n; j ++){
                dp[i][j] += dp[i-1][j];  // 不选第i顶帽子
                dp[i][j] %= mod;
                for(int k = 0; k < n; k ++){   // 选第i顶帽子
                    if((j & (1 << k)) && has[k][i]){  // 枚举所有人，并盼着这个人是不是喜欢第i顶帽子
                        dp[i][j] += dp[i-1][j - (1 << k)];
                        dp[i][j] %= mod;
                    }
                }
            }
        }
        return dp[40][(1 << n) - 1];
    }
};

作者：acw_wangdh15
链接：https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/solution/zhuang-tai-ya-suo-dp-by-wangdh15-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

o*m
shopeemobile.com