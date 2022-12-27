#include <string>
class Solution {
public:
    

    int countDigitOne(int n) {
        //自己不会看的网上的代码和思路
        if(n == 0)return 0;
        string s(to_string(n));
        int h = s[0] - '0';
        int num = pow(10, s.size()-1);
        
        if(n<10 && n>=1)
            return 1;
        if(h == 1)
        {
            return  countDigitOne(num-1) + (n-h*num) +1 + countDigitOne(n-num);
            //0-999 1000-1234
        }else{
            return (h-1)*countDigitOne(num-1) + countDigitOne(num-1) + num +  countDigitOne(n-h*num) ; 
            //0-999 1000-1999 2000-2999 3000-3999 4000-4234
        }
        

    }
};
