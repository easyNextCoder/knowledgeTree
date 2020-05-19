#include <iostream>
#include <unordered_map>
#include <map>
using namespace std; 

class Solution {
public:
    unordered_map<char, int>cp_con;
    int lengthOfLongestSubstring(string s) {
        int tmp_max = 0;
        int first = 0;
        int second = 1;
        for(int i = 0; i<s.size(); i++)
        {
            if( cp_con.count(s[i]) > 0)
            {
                
                if(tmp_max < cp_con.size())
                {
                    tmp_max = cp_con.size();
                }
                unordered_map<char, int>::iterator iter = cp_con.begin();
                int bindex = cp_con[s[i]];
                while(iter!= cp_con.end())
                {
                    if(iter->second < bindex)
                    {
                       cp_con.erase(iter++->first);//map的连续删除方式 
                       
                    }else{
                    	++iter;
					}
                    
            		cout<<"iter";
                }
                cp_con.insert({s[i], i});
                
            }else{
                cp_con[s[i]] = i;
            }
        }
        return tmp_max>cp_con.size()?tmp_max:cp_con.size();
    }
};

int main()
{

	Solution solution;
	cout<<solution.lengthOfLongestSubstring("abcdefadbdfeadf");
	

	  
    return 0;

}
