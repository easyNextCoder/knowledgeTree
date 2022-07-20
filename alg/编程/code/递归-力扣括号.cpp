#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        backtrack(res,n,0,"");
        return res;
    }
    
    //���ݣ�left��ʾ��ʹ�õ�����������'(',right��ʾ��ʹ�õ���������')'
    void backtrack(vector<string>& res,int left,int right,string track){
        if(!right&&!left)res.push_back(track);
        else{
            /*ʹ��һ�������ţ�ͬʱ��ʹ������������1�������ɱ���������Ч����*/
            if(left>0)backtrack(res,left-1,right+1,track+'(');
            /*��ʹ�õ�������������0������������ԭ����������*/
            if(right>0)backtrack(res,left,right-1,track+')');
        }
    }
};


int main()
{
	Solution solution;
	vector<string> rvs = solution.generateParenthesis(3);
	for(auto item: rvs)
		cout<<item<<endl;
	
}
