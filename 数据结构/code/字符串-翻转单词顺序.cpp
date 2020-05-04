#include <iostream>
#include <string>
#include <vector>
#include <algorithm> 
using namespace std;

class Solution {
public:
    string reverseWords(string s) {
    	
        
        string rval;
        int count = 0;
        for(auto item:s){
            if(item == ' ')count++;
        }
        if(count == s.size())return string();
        
        int before = -1;
        int last = -1;
       
       
	    vector<string> svec;
        for(int i = 0; i<s.size(); i++){
        	
        	
            if(s[i] == ' '){
                if(before  !=  -1){
                    svec.push_back(s.substr(before, last-before+1));
                    before = -1;
                }
            }else{
                if(before == -1){
                    before = i;
                    last = before;
                    if(i == s.size()-1){
                        last++;
                        svec.push_back(s.substr(before, last-before+1));
                        break;
                    }
                }else{
                    last++;
                    if(i == s.size()-1){
                        last++;
                        svec.push_back(s.substr(before, last-before+1));
                        break;
                    }
                }
            }
            
        }
        reverse(svec.begin(), svec.end());
        
        rval+=svec.front();
        for(auto iter = svec.begin()+1; iter != svec.end(); iter++){
            rval+=" ";
            rval+=*iter;
        }
        return rval;
    }
};

int main(){
	
	Solution s;
	cout<<s.reverseWords("a");
}
