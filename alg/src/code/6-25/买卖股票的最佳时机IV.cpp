#include <iostream>
#include <vector>

using namespace std;
class Solution {
public:
    vector<int> cprices;
    int K;
    //返回值是tok最大为k的获利最大值
    int f(int tok, int index, int have){

        if(tok == K || index == cprices.size())return 0;//持有到最后还没有卖，就没有意义了
        int maxOut = 0;
        //决定是买还是卖的问题
        if(have)
        {
            int rval1 = 0;
            int rval2 = 0;
            //继续持有c
            
            //[1][2][1] = [2][3][0]+cprices[2];

                        //0 + cprices[2];
                        
            rval1 = f(tok, index+1, have);
            //卖出                                  
            rval2 = f(tok+1, index+1, !have)+cprices[index];
            maxOut = max(maxOut, max(rval1, rval2));
        }else{
            //现在决定是买还是不买的问题
            int rval1 = 0;
            //不买
            rval1 = f(tok, index+1, have);
            //买了
            int rval2 = f(tok, index+1, !have)-cprices[index];
            maxOut = max(maxOut, max(rval1, rval2));
        }
        return maxOut;

    }
    int maxProfit(int k, vector<int>& prices) {

        if(prices.size()<=1 || k == 0)return 0;
        if(k>=prices.size()*2)
        {
            //k过大退化成贪心算法-买卖股票II（看的答案）
            int outCnt = 0;
            for(int i = 1; i<prices.size(); i++)
            {
                if(prices[i]>prices[i-1])
                {
                    outCnt += prices[i]-prices[i-1];
                }
            }
            return outCnt;
        }

        vector<vector<vector<int>>> f(prices.size()+1, vector<vector<int>>(k+2, vector<int>(3, 0)));

        //处理边界条件的都是自己看的题解
        for(int i = 0; i<prices.size(); i++)
        {
            //base case看的答案
            f[i][0][0] = 0;//至今为止没有交易，收益为0
            f[i][0][1] = 0;//交易了0次，但持有股票，不符合规则
            for(int j = 1; j<=k; j++)
            {  
                // base case看的答案
                if(i==0){
                    f[0][j][0] = 0;//第一天买入t次，当天卖出t次,收入为0
                    f[0][j][1] = -prices[i];//甭管第一天买多少次，反正最后少卖一次，持有了股票
                    continue;
                }
                //此时我们手中有股票
                f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0]-prices[i]);//(不卖手中股票， 卖掉手中股票)
            
                //此时我们手中无股票
                //f[1][0][1] f[1][0][0]
                f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1]+prices[i]);
            }
        }             
        return max(f[prices.size()-1][k][0], f[prices.size()-1][k][1]); 
    }
};
int main()
{
    Solution so;
    int k = 2;
    vector<int> vec = {2,4,1};
    cout<<so.maxProfit(k, vec)<<endl;;
}


