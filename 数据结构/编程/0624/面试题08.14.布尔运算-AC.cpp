
//speed up;
int dp[300][300][3] = {0};
class Solution {
public:
    //pair<first, second>first==0 second ==1
    string str;
    pair<int, int> helper(int first, int last)
    {

        if(dp[first][last][2])return{dp[first][last][0], dp[first][last][1]};
        pair<int, int> out = {0,0};

        if(first == last)
        {
            if(str[first] == '0')
                out.first = 1;
            else if(str[first] == '1')
                out.second = 1;
            dp[first][last][2] = 1;
            dp[first][last][0] = out.first;
            dp[first][last][1] = out.second;
            return out;
        }

        for(int i = first; i<=last; i++)
        {
            pair<int, int> tmp = {0,0};
            if(str[i] != '0' && str[i] != '1')
            {
                pair<int, int> left = helper(first, i-1);
                pair<int, int> right = helper(i+1, last);

                switch(str[i])
                {
                    case '^':
                        //生成==1的情况，和==0的情况之后返回
                        tmp.first = left.first*right.first+left.second*right.second;
                        tmp.second = left.first*right.second+left.second*right.first;
                        break;
                    case '|':

                        tmp.first = left.first*right.first;
                        tmp.second = left.first*right.second+left.second*right.first+left.second*right.second;
                        break;
                    case '&':
                        tmp.first = left.first*right.first+left.first*right.second+left.second*right.first;
                        tmp.second = left.second*right.second;
                        break;
                }
            }
            out.first+=tmp.first;
            out.second+=tmp.second;
        }
        dp[first][last][2] = 1;
        dp[first][last][0] = out.first;
        dp[first][last][1] = out.second;
        return out;
    }
	int countEval(string s, int result) {
        memset(dp, 0, sizeof(dp));
		str = s;
        pair<int, int> res = helper(0, str.size()-1);

        if(result == 0)
            return res.first;
        else 
            return res.second;
		return 0;
	}
};