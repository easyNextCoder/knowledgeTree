/*
const int N = 20010;
int f[N][N];//从下标i到j可以获得的最大点数，f[0][nums.size()-1]就是最终的结果

class Solution {
public:
	int deleteAndEarn(vector<int>& nums) {
		if (nums.size() == 0)return 0;
		sort(nums.begin(), nums.end());
		for (int i = 0; i < nums.size(); i++)f[i][i] = nums[i];
		for (int i = 0; i < nums.size() - 1; i++)
		{
			//cout << i << endl;
			for (int len = 1; i + len < nums.size(); len++)
			{
				int j = i + len;
				for (int k = i; k <= j; k++)
				{
					if (k + 1 <= j && nums[k] == nums[k + 1])continue;
					int before = nums[k] - 1;
					int after = nums[k] + 1;
					int kl = k - 1;
					int kr = k + 1;
					int middleValue = nums[k];
					while (kl >= i && (nums[kl] == before || nums[kl] == nums[k]))
					{
						if (nums[kl] == nums[k])middleValue += nums[k];
						kl--;
					}
					while (kr <= j && (nums[kr] == after || nums[kr] == nums[k]))
					{
						if (nums[kr] == nums[k])middleValue += nums[k];
						kr++;
					}

					int maxl = 0;
					int maxr = 0;
					if (kl >= i)maxl = f[i][kl];
					if (kr <= j)maxr = f[kr][j];

					f[i][j] = max(f[i][j], middleValue + maxl + maxr);
				}
			}
		}

		int out = f[0][nums.size() - 1];
		for (int i = 0; i < nums.size(); i++)
		{
			for (int j = 0; j < nums.size(); j++)
				f[i][j] = 0;
		}
		return out;
	}
};
*/

//看了题解中的第2个，变形版的打家劫舍,
//其实这个题应该这样理解从任何一点开始
//寻找最大的结果跟从左向右顺序寻找是一样的
//所以就每有必要用到区间dp的那个算法
//自己一开始使用的就是区间dp的方法，能够
//正确的算出来但是遇到数量是几百的时候就超时

const int N = 20010;
int f[N];
int dp[N][2];//将所有的序列变成打家劫舍的格式（不该写成dp不改了）
int dpc[N];  //真正记录信息的dp数组

class Solution {
public:
	int deleteAndEarn(vector<int>& nums) {
		if (nums.size() == 0)return 0;
		sort(nums.begin(), nums.end());
        int start = nums[0];
        dp[0][0] = start;
        dp[0][1] = 1;
        int countdp = 0;
        for(int i = 1; i<nums.size(); i++)
        {
            if(nums[i] == start)
            {
                dp[countdp][1]++;
            }else{
                countdp++;
                dp[countdp][0] = nums[i];
                start = nums[i];
                dp[countdp][1] = 1;
            }
        }
        dpc[0] = dp[0][0]*dp[0][1];
        if(abs(dp[0][0]-dp[1][0]) != 1)
            dpc[1] = dp[1][0]*dp[1][1]+dp[0][0]*dp[0][1];
        else
            dpc[1] = max(dp[1][0]*dp[1][1], dp[0][0]*dp[0][1]);
        for(int i = 2; i<=countdp; i++)
        {
            int max1 = 0;
            if(abs(dp[i][0]-dp[i-1][0]) != 1)
                dpc[i] = dp[i][0]*dp[i][1] + dpc[i-1];
            else
                dpc[i] = max(dpc[i-2]+dp[i][0]*dp[i][1], dpc[i-1]);
        }

		return dpc[countdp];
	}
};