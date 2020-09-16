class Solution {
public:
    //动态规划-》全练
    //栈-》
    int odesiredTotal = 0;
    int omaxChoosableInteger = 0;
    map<int, bool>dp;
    bool canWin(int tmp, int state)
    {
        if(dp.count(state))return dp[state];
        for(int i = 1; i<=omaxChoosableInteger; i++)
        {
            int cur = 1<<(i);
            if(!(cur & state))
            {
                if(i + tmp>=odesiredTotal || !(canWin(tmp+i, state | cur)))
                {
                    dp[state] = true;
                    return true;
                }
            }
        }
        dp[state] = false;
        return false;
    }

    bool canIWin(int maxChoosableInteger, int desiredTotal) {
        odesiredTotal = desiredTotal;
        omaxChoosableInteger = maxChoosableInteger;
        if(maxChoosableInteger >= desiredTotal)return true;
        else if((1+maxChoosableInteger)*maxChoosableInteger/2 <desiredTotal)return false;
        //cout<<"into canWin()"<<endl;
        return canWin(0, 0);
    }
};