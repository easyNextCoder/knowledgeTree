//https://leetcode-cn.com/problems/unique-binary-search-trees/
class Solution {
public:
    int count = 0;
    int dfs(int first, int last)
    {
        if(last <= first)return 1;
        
        int root = first;
        int tmp = 0;
        for(root; root<=last; root++)
        {
            tmp += dfs(first, root-1)*dfs(root+1, last);
        }
        return tmp;
    }

    int numTrees(int n) {
        int *f = new int[n+1];
        f[0] = 1;
        f[1] = 1;
        if(2<n+1)
            f[2] = 2;
        for(int i = 3; i<n+1;i++)f[i]=0;
        for(int i = 3; i<=n; i++)
        {
            for(int j = 0; j<i; j++)
            {
                f[i]+= f[j]*f[i-j-1];  
            }
        }
        return f[n];
    }
};