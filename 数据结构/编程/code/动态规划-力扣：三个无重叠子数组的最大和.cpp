#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	vector<int> maxSumOfThreeSubarrays(vector<int>& nums, int K) { // 注意原题这里是小k，为了和题解中的符号一致，输入改为大K
	
		int N = nums.size();
		int sum = 0;
		vector<int> W(N-K+1);
		for(int i = 0; i<N; i++)
		{
			sum+=nums[i];
			if(i>=K)
			{
				sum -= nums[i-K];	
			}
			
			if(i>=K-1)
			{
				W[i-K+1] = sum;	
			} 
		}
		
		
		
		vector<int> left(W.size(),0);
		int best = 0;
		left[0] = 0;
		for(int i = 1; i<left.size(); i++)
		{
			if(W[best] < W[i])
			{
				best = i;
			}
			left[i] = best;
		}
		
		vector<int> right(W.size(),0);
		best = right.size()-1;
		for(int i = right.size() -1; i>=0; i--)
		{
			if(W[best] <= W[i])
			{
				best = i;
			}
			right[i] = best;
			
		}
		
		vector<int> result(W.size(),0);
		
		int tmp_max = result[K];
		int tmp_max_index = K;
		for(int i = K; i<W.size()-K; i++)//注意i<W.size()-K,整个for循环的控制条件 
		{
			result[i] = W[i] + W[left[i-K]]+W[right[i+K]];
			if(tmp_max < result[i]){
				tmp_max = result[i];
				tmp_max_index = i;
			}
		}
	
		return {left[tmp_max_index-K], tmp_max_index, right[tmp_max_index+K]};
	
	}

}; 
int main()
{
	Solution solution;
	vector<int> vec = {1,2,1,2,6,7,5,1};
	auto rval = solution.maxSumOfThreeSubarrays(vec, 2);
	
	for(auto item:rval)
	{
		cout<<item<<endl;
	}
	 
	 
	return 0;
}


	/*
	    int N = nums.size();
	    vector<int> W(N - K + 1, 0);
	
	    int sum = 0;
	    for (int i = 0; i < N; ++i) {
	        sum += nums[i];
	        if (i >= K) { sum -= nums[i - K]; }
	        if (i >= K - 1) { W[i - K + 1] = sum; }
	    }
	 
	    vector<int> left(W.size(), 0);
	    int best = 0;
	    for (int i = 0; i < W.size(); ++i) {
	        if (W[i] > W[best]) { best = i; } // 注意这里是 >，为了输出是字典序列
	        left[i] = best;
	    }
	
	    vector<int> right(W.size(), 0);
	    best = W.size() - 1;
	    for (int i = W.size() - 1; i >= 0; --i) {
	        if (W[i] >= W[best]) { best = i; } // 注意这里是 >=，为了输出是字典序列
	        right[i] = best;
	    }
	
	    vector<int> ans{-1, -1, -1};
	    for (int j = K; j < W.size() - K; ++j) {
	        int i = left[j - K], k = right[j + K];
	        if (ans[0] == -1 || W[i] + W[j] + W[k] > 
	                W[ans[0]] + W[ans[1]] + W[ans[2]]) {
	            ans[0] = i;
	            ans[1] = j;
	            ans[2] = k;
	        }
	    }
	    return ans;
	*/
