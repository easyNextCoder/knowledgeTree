#include <iostream>
#include <vector>
#include <map>

using namespace std;
//实现排序
class Solution {
public:
    string minWindow(string s, string t) {
        map<char, int> scon;
        map<char, int> tcon;
        for(auto item:t)
        {
            tcon[item]++;
        }
        int left = 0;
        int right = left;
        int count = 0;
        int tsize = tcon.size();
        vector<pair<int, int>> tmpResult;
        while(right < s.size())
        {
            char a = s[right];
            if(++scon[a] == tcon[a])
            {
                count++;
            }
            if(count == tsize)
            {
                tmpResult.push_back(make_pair(left, right));
                while(left<right)
                {
                    char b = s[left];

                    if(tcon.count(b)>0 && --scon[b] < tcon[b])
                    {
                        count--;
                        tmpResult.push_back(make_pair(left, right));
                        left++;
                        ++right;
                        break;
                    }else{
                        left++;
                    }
                }
            }else{
                ++right;
            }
        }
        int tmpMin = INT_MAX;
        int minl = -1;
        int minr = -1;
        for(auto item:tmpResult)
        {
            if(item.second - item.first < tmpMin)
            {
                tmpMin = item.second-item.first;
                minl = item.first;
                minr = item.second;
            }
        }

        return s.substr(minl, tmpMin+1);
    }
};

int main()
{
    Solution solution;
    cout<<solution.minWindow("ADOBECODEBANC","ABC");
    return 0;
}