#include <iostream>
#include <string>
#include <stack>

using namespace std;
 
class Solution {
public:
    double myPow(double x, int n) {
        if(x==0.0)return 0;
        if(n == 0)return 1;

        long long int nn = n;
        
        double table[32] = {0};
        table[0] = x;
        for(int i = 1; i<32; i++)
        {
            table[i] = table[i-1]*table[i-1];
        }

        long long int posn = nn>0?nn:-nn;
        double result = 1;
        long long int i = 31;

        while(posn){
            long long int step = (long long int)((long long int)1<<i);
            if(posn - step < 0)
            {
                i--;
                cout<<i<<endl;
            }else if(posn-step >= 0)
            {
                result *= table[i];
                posn -= step;
                
            }
        }
        
        return nn<0?(1/result):result;
    }
};

int main()
{
	Solution solution;
	cout<<solution.myPow(2.0,22)<<endl;

	
	return 0;
}


