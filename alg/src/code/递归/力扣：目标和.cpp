//https://leetcode-cn.com/problems/target-sum/
class Solution {
public:
    int count = 0;
    void dfs(int u, int n, vector<int>& nums, int tmpSum, int S)
    {
        if(u == n)
        {
            if(tmpSum == S)count++;
            return;//只是从子叶结点返回！
        }

        dfs(u+1, n, nums, tmpSum-nums[u], S);
        dfs(u+1, n, nums, tmpSum+nums[u], S);
    }
    int findTargetSumWays(vector<int>& nums, int S) {
        dfs(0, nums.size(), nums, 0, S);
        return count;
    }
};