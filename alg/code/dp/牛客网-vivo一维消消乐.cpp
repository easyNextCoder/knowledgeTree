#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <vector>
#include <algorithm>

using namespace std;

/**
 * Welcome to vivo !
 */

#define MAX_NUM 100

void dfs(int boxs[], int N, vector<int> & result, int i, int j, int tmp_res, int last_value)
{
	if (!(i <= j)) {
		result.push_back(tmp_res);
	}
	else {
		int v = 0;
		while (i <= j)
		{
			if (boxs[i] == last_value)
			{
				v++;
				i++;
			}
			else if (boxs[j] == last_value)
			{
				v++;
				j--;
			}
			else {

				dfs(boxs, N, result, i, j, tmp_res + v * v, boxs[i]);
				dfs(boxs, N, result, i, j, tmp_res+v*v, boxs[j]);
			}
		}
		result.push_back(tmp_res + v * v);
	}
}
int solution(int boxs[], int N)
{
	// TODO Write your code here
	vector<int> result;
	int tmp_res = 0;
	dfs(boxs, N, result, 0, N - 1, tmp_res, boxs[0]);
	dfs(boxs, N, result, 0, N-1, tmp_res, boxs[N-1]);
	sort(result.begin(), result.end());
	return result.back();
}

int main()
{
	string str("");
	getline(cin, str);
	int boxs[MAX_NUM];
	int i = 0;
	char* p;
	int count = 0;

	const char* strs = str.c_str();
	p = strtok((char*)strs, " ");
	while (p)
	{
		boxs[i] = atoi(p);
		count++;
		p = strtok(NULL, " ");
		i++;
		if (i >= MAX_NUM)
			break;
	}

	int num = solution(boxs, count);
	cout << num << endl;
	return 0;
}


//自己的理解

（1）自己的解法如下图采用
for()
{
	递归函数()
}
的方法的时间复杂度n!无法计算大于15的结果
只能转化为记忆化搜索的方式，dp来解答

class Solution {
public:
    int removeBoxes(vector<int>& boxes) {
        if(boxes.empty())return 0;
        
        vector<pair<int, int>> shrink;
        shrink.push_back({boxes[0], 1});
        int start = boxes[0];
        for(int i = 1; i<boxes.size(); i++)
        {
            if(boxes[i] == start)
            {
                shrink.back().second++;
            }else{
                start = boxes[i];
                shrink.push_back({start, 1});
            }
        }
        //这种本来可以使用区间dp的方式，但是由于每次选择一点的时候
        //左右区间还有关系，无法使用简单的二维dp，所以只能分情况讨论
        //选择当中最大的那个，而且每次选择区间被一分为二，
			 //但是递归中有for循环所以无法降低时间复杂度
        //dp解决的问题有一个明显的特点就是，有重复子，对于这个题如何找到合适的重复子？
        int out = 0;
        int leftm = dfs(1, shrink.size()-1)+shrink[0].second*shrink[0].second;
        int rightm = dfs(0, shrink.size()-2)+shrink[shrink.size()-1].second*shrink[shrink.size()-1].second;


        for(int i = 1; i<shrink.size()-1; i++)
        {
            int k = shrink[i].second*shrink[i].second;
            int tmp_max = 0;
            if(shrink[i-1].first == shrink[i+1].first)
            {
                //将i+1位置的个数移动到i-1位置，移动后再相反的移动
						  //计算二者的最大值，肯定要比不移动计算的结果大
                int tmp = shrink[i+1].second;
                shrink[i-1].second+=tmp;
                int left = dfs(0, i-1, shrink) + shrink[i].second*shrink[i].second+ dfs(i+2, shrink.size()-1,shrink);
                int tmp1 = shrink[i-1].second;
                shrink[i+1].second = tmp1;
                int right = ;
                
            }else{


            }
        }
    }
};

（2）照抄题解并看懂了
https://leetcode-cn.com/problems/remove-boxes/
int dp[100][100][100];
class Solution {
public:
    
    int calPoints(vector<int>& boxes, int dp[][100][100], int l, int r, int k)
    {
        if(l>r)return 0;
        if(dp[l][r][k] != 0)return dp[l][r][k];
			 //这就是记忆搜索了，降低时间复杂度靠的就是它
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
                dp[l][r][k] = max(dp[l][r][k], calPoints(boxes, dp, l, i, k+1)+calPoints(boxes, dp, i+1, r-1, 0));
            }
        }


        return dp[l][r][k];
    }
    int removeBoxes(vector<int>& boxes) {
        memset(dp, 0, sizeof(dp));
        return calPoints(boxes, dp, 0, boxes.size()-1, 0);
    }
};

