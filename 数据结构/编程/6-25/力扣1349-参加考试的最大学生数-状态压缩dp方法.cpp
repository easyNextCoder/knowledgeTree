class Solution {
public:
    int maxStudents(vector<vector<char>>& seats) {
        //状态压缩dp
        int width = seats.front().size();
        vector<vector<int>> dp(seats.size()+1, vector<int>(1<<seats.front().size(), 0));

        vector<int> compressed(seats.size(), 0);
        for(int i = 0; i<seats.size(); ++i)
        {
            int ans = 0;
            for(int j = 0; j<seats[i].size(); ++j)
                ans |= (int)(seats[i][j] == '.')<<j;
            compressed[i] = ans;
        }

        int ans = 0;

        for(int i = 0; i<seats.size(); ++i)
        {//对于每一排
            for(int j = 0; j < (1<<width); ++j)//对于每一排的安排方式
            {
                if( ( j & ~(compressed[i]) ) > 0 )continue;//有的人没有板凳
                if((j & (j<<1)) > 0 || (j&(j>>1)) > 0)continue;//相邻座位上不能安排人
                int stuNum = 0;
                for(int s = 0; s<width; ++s)
                    if(j&(1<<s))
                        stuNum++;//统计当前行坐了多少人

                if(i == 0){//第一排不用考虑前面一排的情况
                    dp[i][j] = stuNum;
                    continue;
                }else{//要考虑前一排的情况
                    for(int k = 0; k<(1<<width); ++k)
                    {
                        if( (k & (j<<1)) > 0 || (k&(j>>1)) > 0 )//上下行不符合规则
                            continue;
                        else  
                            dp[i][j] =  max(dp[i][j], dp[i-1][k]+stuNum);
                            //从上一行转移到下一行,要加上当前行坐的人数
                    }
                }
                ans = max(ans, dp[i][j]);
            }
        }
        return ans;
    }
};


/*
    [["#",".","#","#",".","#"],[".","#","#","#","#","."],["#",".","#","#",".","#"]]
[["#",".",".",".","#"],[".","#",".","#","."],[".",".","#",".","."],[".","#",".","#","."],["#",".",".",".","#"]]


 */
