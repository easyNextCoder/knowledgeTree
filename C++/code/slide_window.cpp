#include<iostream>
#include<vector>
#include<string>
#include <vector>

using namespace std;
 
class Solution {
public:
    vector<int> maxInWindows(const vector<int>& num, unsigned int size)
    {
       if(size == 0 || size>num.size()){
           return vector<int>();
       }
        vector<int>dp;
        int tmp_max = num[0];
        for(size_t j =0; j<size; j++){
            if(tmp_max<num[j]){
                tmp_max = num[j];
            }
        }
        dp.push_back(tmp_max);
        for(size_t i = 1; i<num.size()-size+1; i++){
            if(num[i-1] == dp[i-1]){
                //重新找最大值
                int tmp_max = num[i];
                for(size_t j = i; j<i+size; j++){
                    if(tmp_max<num[j]){
                        tmp_max = num[j];
                    }
                }
                dp.push_back(tmp_max);
            }else if(num[i+size-1]>dp[i-1]){
                dp.push_back(num[i+size-1]);
            }else{
                dp.push_back(dp[i-1]);
            }
        }
        
        return dp;
    }
};

int main(){
	Solution* solution = new Solution();
	vector<int> input = {2,3,4,2,6,2,5,1};
	vector<int> rval = solution->maxInWindows(input,3);
	for(auto item:rval){
		cout<<item<<endl;
	}
	return 0;
}
