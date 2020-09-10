// 使用记忆搜索太难了：因为[l, r]并不能直接表示最大分数这个分数还是依赖于之前的移动对当前数组的影响。
//                   为了保存这个信息，我们需要在记忆化数组中加上额外的一维，告诉当前子序列有多少
//                   个元素被合并在一起。

#include <iostream>
#include <vector>

using namespace std;

int  dp[100][100][100];



class Solution {
public:
    int calPoints(vector<int>& boxes, int dp[][100], int l, int r, int k)
    {
        if(l>r)return 0;
        if(dp[l][r][k] != 0)return dp[l][r][k];//这就是记忆搜索了
        while(r>l && boxes[r] == boxes[r-1])
        {
            r--;
            k++;
        }
        dp[l][r][k] = calPoints(boxes, dp, l, r-1, 0) + (k+1)*(k+1);
        for(int i = l; i<r; i++)
        {
            if(boxes[i] == boxes[r])
            {
                dp[l][r][k] = max(dp[l][r][k], calPoints(boxes, l, i, k+1)+calPoints(boxes, dp, i+1, r-1, 0));
            }
        }
    }
    int removeBoxes(vector<int>& boxes) {
        
        return calPoints(boxes, dp, 0, boxes.size()-1, 0);
    }
};

int main()
{
    vector<int> box = {1,3,2,2,2,3,4,3,1};
    
    Solution so;
    so.removeBoxes(box);


    return 0;
}