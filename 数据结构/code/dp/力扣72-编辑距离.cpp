class Solution {
public:
    /*
    int count = 0;
    vector<int> rval;
    void match(string a, string b, int u)
    {   
        if(a.empty())
        {
            rval.push_back(u+b.size());
            return;
        }else{
            if(b.empty())
            {
                rval.push_back(u+a.size());
                return;
            }
        }
        if(a == b)
        {
            rval.push_back(u);
            return ;
        }
        

        if(a[0] == b[0])
        {
           match(a.substr(1), b.substr(1), u); 
        }else{
            match(a, b.substr(1), u+1);
            //insert
            match(a.substr(1), b, u+1);
            //del
            match(a.substr(1), b.substr(1), u+1);
            //replace
        }
    }
    */
    int minDistance(string word1, string word2) {
        vector<vector<int>> dp(word1.size()+1, vector<int>(word2.size()+1, 0));
        for(int i = 0; i<=word1.size(); i++)
            dp[i][0] = i;
        for(int j = 0; j<=word2.size(); j++)
            dp[0][j] = j;

        for(int i = 1; i<=word1.size(); i++)
        {
            for(int j = 1; j<=word2.size(); j++)
            {
                if(word1[i-1] == word2[j-1])
                    dp[i][j] = dp[i-1][j-1];
                else{
                    dp[i][j] = min(min(dp[i-1][j-1]+1, dp[i-1][j]+1), dp[i][j-1]+1);
                    //把上面一行写成了min(min(dp[i-1][j-1]+1, dp[i-1][j])+1, dp[i][j-1]+1);
                    //这个bug真可笑
                }
                    
            }
        }

        return dp[word1.size()][word2.size()];
    }
};
