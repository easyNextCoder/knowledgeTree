class Solution {
public:
    int last = 0;
    int maxTimes = 2;
    vector<int>cprices;
    int getMax(int timesNow, int first)//(已经购买的次数， 当前在第几天)
    {
        if(timesNow == 2 || last-first+1<2)return 0;//购买了超过两次，或者购买到最末尾了返回
        int maxOut = 0;
        maxOut = getMax(timesNow, first+1);//当前不买，不消耗购买次数
        for(int len = 2; first+len-1<=last; len++)//购买股票的区间,有可能购买2到末尾的天数
        {
            maxOut = max(maxOut, getMax(timesNow+1, first+len)+cprices[first+len-1]-cprices[first]);
        }
        return maxOut;
    }

    int maxProfit(vector<int>& prices) {
        if(prices.size()<=1)return 0;
        cprices = prices;
        last = prices.size()-1;
        //求完成两笔交易的最多的获利
        return getMax(0, 0);
    }
};