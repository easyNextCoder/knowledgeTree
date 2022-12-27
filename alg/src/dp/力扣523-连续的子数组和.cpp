//https://leetcode-cn.com/problems/continuous-subarray-sum/

const int N = 10010;
int sum[N];

class Solution {
public:
    
    
    bool checkSubarraySum(vector<int>& nums, int k) {
        if(k == 0){
            for(int i = 0; i<nums.size()-1; i++)
            {
                if(nums[i] == nums[i+1] && nums[i] == 0)
                {
                    return true;
                }
            }
            return false;
        }
        
        //可以 使用滑动窗口? 好像不行，滑动窗口求的是刚好符合条件的值
        memset(sum, 0, sizeof(sum));
        for(int i = 1; i<=nums.size(); i++)
        {
            sum[i] = sum[i-1]+nums[i-1];
        }

        for(int len = 2; len<=nums.size(); len++)
        {
            for(int i = 0; i+len-1<nums.size(); i++)
            {
                int j = i+len-1;
                int tmp = sum[j+1]-sum[i];
                if(tmp%k == 0)return true;
            }
        }

        return false;
    }


};