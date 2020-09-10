int f[35000][3];
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size()<=1)return 0;
        
        memset(f, 0, sizeof(f));
        int minal = prices[0]>prices[1]?prices[1]:prices[0];
        f[0][0] = 0;
        f[0][1] = 0;
        f[1][0] = prices[1]-prices[0]>0?prices[1]-prices[0]:0;
        f[1][1] = f[1][0]>0?f[1][0]:0;
        for(int i = 2; i<(int)prices.size(); i++)
        {//在这里每一步我们进行两次转移，因为一次转移代表着要买这个区间段，在之前的买卖股票问题中我们可以买卖
         //无数次，所以可以进行不间断的转移！思考
               
            f[i][0] = f[i-1][0];//今天不买股票,那么今天就是之前天数能获利的最大值
            if(prices[i]-minal > f[i-1][0])
            {
                f[i][0] = prices[i]-minal;
            }else{
                f[i][0] = f[i-1][0];   
            }

            if(minal >= prices[i])//看的后面的题解说要剪枝
            {
                minal = prices[i];
                f[i][1] = f[i-1][1];
                continue;
            }
              
            //今天买入股票
            f[i][1] = f[i-1][1];//?这一步是自己爆数据debug之后加的[4657,8368,3942,1982,5117,563,3332]
            for(int len = 1; len<=i; len++)//有可能买卖的时间间隔就是1天，也有可能买卖的时间间隔是从0-i
            {
                if(prices[i] > prices[i-len])//在这里进行了剪枝（参照后面的题解）
                    f[i][1] = max(f[i][1], f[i-len-1<0?0:i-len-1][0]+prices[i]-prices[i-len]);
                //在这一步，我们已经买了一次了
            }
            //cout<<f[i][0]<<":"<<f[i][1]<<endl;
        }
        return max(f[prices.size()-1][0], f[prices.size()-1][1]);
    }
};
总结是：递归的形式最终决定了你动态规划的形式！
1.这一题自己的递归和动态规划的效率都是在N^2级别
2.看看别人写的递归的最终形式也最终决定了其动态规划的最后的效率是O(n)级别的复杂度
  https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/comments/28872
3.学会如何剪枝！
	自己在写的过程中一直通不过最后看了题解中的如何剪枝才过的。
	还有一个总结就是：剪枝非常重要！是效率不高的递归改动态规划的救星

