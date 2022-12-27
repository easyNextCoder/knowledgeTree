const int N = 1050;
bool f[N][N];
class Solution {
public:
    bool isMatch(string s, string p) {
        int m = s.size();
        int n = p.size();
        memset(f, 0, sizeof(f));
        f[0][0] = true;
        for (int i = 0; i <= m; ++i) {
            for (int j = 1; j <= n; ++j) {
                if (p[j - 1] == '*') {
                     
                    //匹配0次
                    f[i][j] |= f[i][j-1];
                    //或者匹配1次，就像递归
                    if(i>0)
                        f[i][j] |= f[i-1][j];
                    //这个根本看起来就像是一个自带优化的递归
                    //当你把f看成一个函数，它有两个参数，
                    //每进行过一次计算就将结果保存起来供下次调用
                    //跟记忆化递归时使用的函数何尝不是一样的？
                }
                else {
                    if(i == 0)
                    {
                        f[i][j] = false;
                        continue;
                    }
                    if(s[i-1] == p[j-1] || p[j-1] == '?')
                    {
                        f[i][j] = f[i-1][j-1];
                    }
                    
                }
            }
        }
        // for(int i = 0; i<=m; i++)
        // {
        //     for(int j = 0; j<=n; j++)
        //         cout<<f[i][j]<<" ";
        //     cout<<endl;
        // }
        return f[m][n];
    }
};