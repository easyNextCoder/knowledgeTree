#include <iostream>
#include <vector>
#include <algorithm>
#include <bitset>
using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
    	if(prices.size() == 0 || prices.size() == 1){
    		return 0;
		}else{
		
			adjacent_difference(prices.begin(), prices.end(), prices.begin());
			/*
			for_each(prices.begin(), prices.end(),[](int a){
				cout<<a<<endl;
			});
			*/
			vector<int> nprices(++prices.begin(), prices.end());
			//计算连续的最大和
			vector<int>dp(nprices.size());
			int max = nprices[0];
			int smax = nprices[0];
			dp[0] = nprices[0];
			for(int i = 1; i<nprices.size(); i++){
				if(nprices[i] + dp[i-1] > nprices[i]){
					dp[i] = nprices[i] + dp[i-1];
				}else{
					dp[i] = nprices[i];
				}
				if(dp[i]>max){
					max = dp[i];
				}
			} 
			if(max<0)max = 0;
			return max;
		}
    }
};


int main(){
	/*
	Solution solution;
	vector<int> vec = {1,0};
	cout<<solution.maxProfit(vec);
	*/
	cout<<INT_MAX<<endl;
	cout<<INT_MIN<<endl;
	bitset<32>x(-1);
	cout<<x<<endl;
	cout<<x.to_ulong()<<endl;
	bitset<32>y(INT_MIN+INT_MAX);
	cout<<y<<endl;
	cout<<y.to_ulong()<<endl;
	cout<<(int)(INT_MIN+INT_MIN)<<endl;
	cout<<(int)(INT_MAX+INT_MAX)<<endl;
	cout<<((int)(-1-1073741824))<<endl;
	
	//unsigned
	bitset<32>us(-1);
	cout<<us<<endl; 
	cout<<sizeof(size_t)<<endl;
	cout<<sizeof(long)<<endl;
	cout<<sizeof(long long int)<<endl;
	return 0;
}
 
