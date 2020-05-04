#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

/* //直接采用递归的方法时间上通不过 
class Solution {
public:
    void cacul(vector<int>&vec, int n, int totalNum){
        if(n == 0){
			vec[totalNum]++;
			return;
        } 
        for(int i = 0; i<6; i++){
            totalNum += i;
        	cacul(vec, n-1, totalNum);
            totalNum -= i;
        }

    }
    vector<double> twoSum(int n) {
        vector<int>vec(n*6+1,0);
        
        int minPoints = n;
        cacul(vec, n, minPoints);
        
        int out = 0;
        for_each(vec.begin(), vec.end(), [&out](int&a){
			out+=a;
		}) ;
		
        vector<double>dvec;
        for_each(vec.begin(), vec.end(), [&dvec, &out](int & a){
            if(a != 0){
                dvec.push_back((double)a/out);
            }
        });
        
        return dvec;
	}
};
*/

class Solution {
public:
   //这次采用动态规划的方法，方法的思想见剑指offer上 
    vector<double> twoSum(int n) {
        static int maxPoints = 6;
		int vs[2][n*maxPoints+1]; 
		for(int i = 0; i<2; i++){
			for(int j = 0; j<n*maxPoints+1; j++){
				vs[i&1][j] = 0;
			}
		}
		
        //init
		for(int i = 1; i<=maxPoints; i++)
			vs[1&1][i]++;
        
        for(int i = 2; i<=n; i++){
        	//骰子的个数
        	int index = i&1;
			for(int j = 1; j<=n*maxPoints; j++){
				if(vs[1-index][j] != 0){
					for(int k = 1; k<=maxPoints; k++){
						vs[index][j+k] = vs[index][j+k] + vs[1-index][j];
					}
				}
			}
			for(int j = 1; j<=n*maxPoints;j++){
				vs[1-index][j] = 0; 
			}
		}
		
		
		int count = 0;
		for(int i = 1; i<=n*maxPoints; i++){
			count += vs[n&1][i];
		}
		
		vector<double>vec;
		double dcount = (double)count;
		for(int i = 1; i<=n*maxPoints; i++){
			if(vs[n&1][i] != 0)
				vec.push_back((double)vs[n&1][i]/dcount);
		}
		for_each(vec.begin(), vec.end(),[](double&a){
			cout<<a<<endl;
		});
		  
        
        return vec;
	}
};

int main(){
	
	Solution solution;
	solution.twoSum(3);
	
	return 0;
}
