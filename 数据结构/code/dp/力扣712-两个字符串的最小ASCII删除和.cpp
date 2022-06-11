方法1：使用lcs算法
也可以转化为使用最长公共字串来做，只要求得最长公共字串，并且保证最长公共字串的和最大就可以了。

其实就本题而言，从lcs的意义上去推导这个转移公式是更加容易做的，直接使用lcs算法要加辅助去寻找那个最大的ASIC值的字符串

class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        int n1 = s1.length(),n2 = s2.length();
        int sum = 0;
        int[][] dp = new int[n1+1][n2+1];
        for(int i = 0;i < n1;i++){
            sum += s1.charAt(i);
            for(int j = 0;j < n2;j++)
                if(s1.charAt(i) == s2.charAt(j))
                    dp[i+1][j+1] = dp[i][j] + s1.charAt(i); //
									//实际在dp方格中每个相等的点就求出了
									//两字符串截止到当前位置的最长公共子串   
                else 
                    dp[i+1][j+1] = Math.max(dp[i][j+1],dp[i+1][j]);
        }
        for(int i = 0;i < n2;i++) sum += s2.charAt(i);
        return sum - dp[n1][n2]*2;
    }
}

方法2：根据定义来刨析如何进行状态转移（所有的状态转移公式，都是lcs转移公式的变种）
本题状态转移公式：

for(int j = 1; j<=s2.size(); j++)
            {
                if(s1[i-1] == s2[j-1])//一开始是s1[i-1] == s2[j-1]在这里bug了20分钟
                {
                    dp[i][j] = dp[i-1][j-1];
                }else{
                    dp[i][j] = min(dp[i-1][j]+s1[i-1], dp[i][j-1]+s2[j-1]);
                }
            }


