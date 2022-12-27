class Solution {
public:
    vector<int> cprices;
    int ssize = 0;
    int getProfit(int index, int have)
    {
        if(index >= ssize)return 0;
        //index==0而不管have什么事情就决定了，下面的动态规划对第一维为0第二维为任意的情况要先做出给定！
        if(have)
        {
            return max( getProfit(index+1, have), getProfit(index+2, !have)+cprices[index] );       
        }else{
            return max(getProfit(index+1, have), getProfit(index+1, !have)-cprices[index]);
        }
    }
    /*
        cprices = prices;
        ssize = cprices.size();
        return getProfit(0, 0);
     */

    int maxProfit(vector<int>& prices) {
        if(prices.size()<=1)return 0;
        vector<vector<int>> f(prices.size()+1, vector<int>(3, 0));
        //以下的base条件是通过上面的递归式推导出来的
        f[0][0] = 0;
        f[0][1] = 0-prices[0];
        f[1][0] = max(f[0][0], f[0][1]+prices[1]);
        f[1][1] = max(f[0][1], 0-prices[1]); 
        for(int i = 2; i<prices.size(); i++)
        {
            f[i][0] = max(f[i-1][0], f[i-1][1]+prices[i]);
            f[i][1] = max(f[i-1][1], f[i-2][0]-prices[i]);
        }
        return max(f[prices.size()-1][0], f[prices.size()-1][1]);
    }

};