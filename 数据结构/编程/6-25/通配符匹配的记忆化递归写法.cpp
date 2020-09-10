
//使用朴素的动态规划的方式，根本不可能！

const int N = 1101;
bool su[N][N][2];//speed up
class Solution {
public:
    
    bool iisMatch(string s, string p) {

        if(su[s.size()][p.size()][0])return su[s.size()][p.size()][1];
        bool rval0 = false;
        if(s.empty())
        {
            if(p.empty())
            {
                rval0 = true;
            }else{
                if(p[0] == '*')
                {
                    rval0 = iisMatch(s, p.substr(1));
                }else{
                    rval0 = false;
                }
            }
            su[s.size()][p.size()][0] = true;
            su[s.size()][p.size()][1] = rval0;
            return rval0;
        }

        if(p.empty())
        {
            rval0 = s.empty();
            su[s.size()][p.size()][0] = true;
            su[s.size()][p.size()][1] = rval0;
            return rval0;
        }

        if(s[0] == p[0] || p[0] == '?')
        {
            rval0 = iisMatch(s.substr(1), p.substr(1));
        }else if(p[0] == '*')
        {//匹配0次 or 匹配1次
            rval0 = iisMatch(s, p.substr(1)) || iisMatch(s.substr(1), p);
        }

        su[s.size()][p.size()][0] = true;
        su[s.size()][p.size()][1] = rval0;
        return rval0;
    }

    bool isMatch(string s, string p)
    {
        memset(su, 0, sizeof(su));
        return iisMatch(s, p);
    }

};