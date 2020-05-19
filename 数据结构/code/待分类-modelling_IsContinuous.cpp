#include <iostream>
#include <string>
#include <algorithm>
#include <numeric> 
using namespace std;


class Solution {
public:
    bool IsContinuous( vector<int> numbers ) {
        if(numbers.size() != 5)
            return false;
        
        sort(numbers.begin(), numbers.end(), [](int a, int b){
            return a<b;
        });
        
        auto iter = find_if(numbers.begin(), numbers.end(), [](int a){
        	return a != 0;
		});
        	
        	
    	vector<int>tmp(iter, numbers.end());
        int old_tmp_size = tmp.size();
		tmp.erase(unique(tmp.begin(), tmp.end()), tmp.end());
        
		if(tmp.size()<old_tmp_size)
            return false;
        else{
            
            adjacent_difference(tmp.begin(), tmp.end(), tmp.begin());
            tmp[0] = 0;
            int sum = 0;
            for_each(tmp.begin(), tmp.end(),[&sum](int& a){
            	sum+=a;
			});
			
			return sum<5;
        }
    
    }
};

int main(){
	Solution solution;
	cout<<solution.IsContinuous({1,3,0,5,0})<<endl;;
	return 0;
}
