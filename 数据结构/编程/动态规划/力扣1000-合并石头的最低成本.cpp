//https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
//int f[110][110][110];//f[i][j][k] 将i到j区间的石头合并成k块的最小代价


class Solution {
public:
    
    int mergeStones(vector<int>& stones, int K) {
        int sum[110];
        int n = stones.size();
        if((n-1)%(K-1) != 0)return -1;

        //求最小代价是应该提前将所有dp内容设置为极大值
        //相反求最大代价应该是提前将所有dp内容设置为极小值
        vector<vector<vector<int>>> f(n, vector<vector<int>>(n, vector<int>(K + 1, 1e8)));

        for(int i = 0; i<n; i++)
        {
            sum[i+1] = sum[i]+stones[i];
        }

        //这里先初始化f[i][i][1]就是为了下面的k=2的for循环做铺垫
        for(int i = 0; i<n; ++i)
            f[i][i][1] = 0;

        //这里子问题的长度为什么又要从2开始？
        //就是从2开始，因为len控制着从页到根的迭代过程
        //不能出现违法的叶子
        for(int len = 2; len<=n; len++)
        {
            for(int i = 0; i+len-1<n; i++)
            {
                int j = i+len-1;
                //给定了区间之后，开始进行分堆
                for(int k = 2; k<=K; k++)
                {
                    //这里p从i开始就是为了利用前文初始化的f[i][i][1] = 0;
                    //!注意这里的p+=k-1
                    for(int p = i; p<j; p+=K-1)
                    {
                        f[i][j][k] = min(f[i][j][k], f[i][p][1]+f[p+1][j][k-1]);
                        //将k个堆为1堆和k-1堆, 而f[p+1][j][k-1]又是将区间[p+1,j]划分为k-1堆的最小花费
                    }
                }
                f[i][j][1] = f[i][j][K] + sum[j+1]-sum[i];
                //以上两个for循环实际作用替代了一个(~j^k)的递归时间复杂度过高（复杂度计算？）
            }
        }

        return f[0][n-1][1];
        
        
    }
};
/*
//https://zxi.mytechroad.com/blog/dynamic-programming/leetcode-1000-minimum-cost-to-merge-stones/
class Solution {
public:
  int mergeStones(vector<int>& stones, int K) {
    const int n = stones.size();
    if ((n - 1) % (K - 1)) return -1;
    const int kInf = 1e9;    
    vector<int> sums(n + 1);
    for (int i = 0; i < stones.size(); ++i)
      sums[i + 1] = sums[i] + stones[i];
    // dp[i][j][k] := min cost to merge subarray i~j into k piles.
    vector<vector<vector<int>>> dp(n, vector<vector<int>>(n, vector<int>(K + 1, kInf)));
    for (int i = 0; i < n; ++i)
      dp[i][i][1] = 0;
    
    for (int l = 2; l <= n; ++l) // subproblem length
      for (int i = 0; i <= n - l; ++i) { // start
        int j = i + l - 1; // end
        for (int k = 2; k <= K; ++k) // piles
          for (int m = i; m < j; m += K - 1) // split point
            dp[i][j][k] = min(dp[i][j][k], dp[i][m][1] + dp[m + 1][j][k - 1]);
        dp[i][j][1] = dp[i][j][K] + sums[j + 1] - sums[i];
      }        
    return dp[0][n - 1][1];
  }
};
*/