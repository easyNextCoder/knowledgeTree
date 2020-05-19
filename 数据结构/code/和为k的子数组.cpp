#include <iostream> 
#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    //采用滑动窗口算法
    int subarraySum(vector<int>& nums, int k) {
    	//最终还是要使用unordered_map 进行前缀和优化 
    	typedef int preSum;
        typedef int times;
        unordered_map<preSum, times> con;
        con[0]++;
        int sum = 0;
        int count = 0;
        for(int i = 0; i<nums.size(); i++)
        {
            sum+=nums[i];
            if(con[sum-k]>0)
            {
                count+=con[sum-k];
            }
            con[sum]++;
        }
        return count;
    	/*
        int result = 0;
        if(nums.size() == 1)
        {
            if(nums[0] == k)return 1;
        }else{
            int first = 0;
            int last = 0;

            int sum = 0;
            sum += nums[first];
            while(!(last == nums.size()-1 && first == last))
            {
            	cout<<"first last:"<<first<<last<<endl;
            	cout<<"sum k:"<<sum<<k<<endl;
                if(sum == k )
                {
                    result++;
                    if(last < nums.size())
                    {
                        sum+=nums[++last];
                    }
                    else{
                        sum-=nums[first++];
                    }
                }
                else if(sum > k){
                    
                    if(last+1 < nums.size() && nums[last+1]<=0)
                    {
                        sum += nums[++last];
                    }else if(first<last && nums[first]>=0){
                        sum -= nums[first++];
                    }else {//if(first == last)还能处理上面两个条件都不满足的情况
                    	if(last<nums.size())
                    		last++;
                    	if(first < last)
                    		first++;
                    }
                    
                }else if(sum < k){
                    if(last+1 < nums.size() && nums[last+1]>=0)
                    {
                        sum += nums[++last];
                    }else if(first<last && nums[first]<=0){
                        sum -= nums[first++];
                    }else{//if(first == last)
                        if(last<nums.size())
                    		last++;
                    	if(first < last)
                    		first++;
                    }
                }    
            }
			if(sum == k)result++;   
        }
        return result;
        */
    }
};

int main()
{
	Solution solution;
	vector<int> vec = {1,0,0,2,0,6,0,0,8,8,0}; 
	
	/*
	[1,2,3]
	3
	*/ 
	cout<<solution.subarraySum(vec, 8);
	
	return 0;
}
