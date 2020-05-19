#include <iostream>
#include <vector>
#include <random>

using namespace std;
/*
int  main()
{
	default_random_engine e;

	vector<int> vec;
	
	for(int i = 0; i<100; i++)
	{
		cout<<vec.push_back(e())<<endl;
	}
	
	//generator stl
	
	int size = vec.size()-1;
	for(int i = 0; i<vec.size(); i+=size/2-1)
	{
		
		size = size/2;
		
	 } 
	return 0;
}

*/

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
    	//递归问题，如果不剪枝，时间复杂度太高
    	/*
        if(amount == 0)return amount;

        int result_min = INT_MAX;
        int tmp = INT_MAX;
        
        //cout<<"start:"<<endl;
        for(int i = coins.size()-1; i>=0; i--)
        {

			if(amount - coins[i] <0)
				return result_min+1;
			else
            	tmp = coinChange(coins, amount-coins[i]);
            
           
            if(tmp >= 0 && tmp < result_min)
            {
                result_min = tmp;
            }
            //cout<<tmp<<endl;
        }
        //cout<<"end."<<endl;

        if(result_min == INT_MAX){
             return -1;
        }
           
        return result_min+1;
        */
        //如果钞票面值过大，无法申请内存
		/* 
		vector<int> table((long long int)coins.back()+1,0);
		vector<int>newCoin = {0};
		newCoin.insert(newCoin.end(), coins.begin(), coins.end());
		
		for(int i = 1; i<=amount; i++)
		{
			int tmp_min = INT_MAX;
			int min = INT_MAX; 
			for(int j = 1; j<newCoin.size(); j++)
			{	
				int money_value = newCoin[j];
				tmp_min = i-money_value < 0||table[(i-money_value)%table.size()] == INT_MAX?INT_MAX:1+table[(i-money_value)%table.size()];
				
			cout<<tmp_min<<endl;
				if(min > tmp_min)
					min = tmp_min;  
			}
			table[i%table.size()] = min;
			if(i == amount)
				if(min == INT_MAX)
					return -1;
				else
					return min;
		}
		return -1;
		*/
		vector<int>table(amount+1, 0);
		for(int i = 1; i<=amount; i++)
		{
			int tmp_min = INT_MAX;
			int min = INT_MAX;
			for(int j = 0; j<coins.size(); j++)
			{
				int money_value = coins[j];
				if(i-money_value >= 0)
				{
					tmp_min = table[i-money_value] == INT_MAX?INT_MAX:table[i-money_value]+1;
				}else{
					tmp_min = INT_MAX;
				}
				if(min > tmp_min)
					min = tmp_min;
			}
			table[i] = min;
			if(i == amount)
				if(min == INT_MAX)
					return -1;
				else
					return min;
		}
    }
};

int main()
{
	Solution solution;
	vector<int> vec = {2147483647};
	cout<<"result is:"<<solution.coinChange(vec ,103);
}
