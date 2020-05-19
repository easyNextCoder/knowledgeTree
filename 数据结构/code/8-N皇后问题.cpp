#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

//record 目前record没有作用 
//https://www.nowcoder.com/practice/fbf428ecb0574236a2a0295e1fa854cb?tpId=61&&tqId=29558&rp=1&ru=/activity/oj&qru=/ta/pku-kaoyan/question-ranking 
void dfs(bool record[][8], vector<int>& result, int startLineN, vector<int>&colRecord)
{
    if(startLineN == 8)
    {
        int outValue = colRecord[0]+1;
        for(int i = 1; i<colRecord.size(); i++)
        {
            outValue*=10;
            outValue+=colRecord[i]+1;
        }
        result.push_back(outValue);
    }else{
        for(int i = 0; i<8; i++)//将startLineN这一行摆到那一列上
        {
            //到这里来选列数
            if( find(colRecord.begin(), colRecord.end(), i) == colRecord.end() )
            {
                int x = i+1, y = startLineN-1;
                int lastIndex = colRecord.size()-1;
                bool pointCanChose = true;
                for(x,y; x<8&&y>=0; ++x, --y)
                {
                    if(colRecord[lastIndex--] == x)
                    {
                        //重复点这个点不要;
                        pointCanChose = false;
                        break;
                    }
                }
                if(pointCanChose)
                {
                    x = i-1;
                    y = startLineN-1;
                    lastIndex = colRecord.size()-1;
                    for(x,y; x>=0&&y>=0; --x, --y)
                    {
                        if(colRecord[lastIndex--] == x)
                        {
                            //重复点这个点不要;
                            pointCanChose = false;
                            break;
                        }
                    }
                }
                if(pointCanChose == false)
                {
                    continue;
                }else{
                    colRecord.push_back(i);
                    dfs(record, result, startLineN+1, colRecord);
                    colRecord.pop_back();
                }
            }
        }
    }
}

int main()
{
    bool record[8][8] = {0};
    vector<int>result;
    int startLineN = 0;
    vector<int> colRecord;
    dfs(record, result, startLineN, colRecord);
    sort(result.begin(), result.end());
    int n = 0;
    //cin>>n;
    //cin.ignore();
    int tmpIndex = 0;
    while(cin>>tmpIndex)//cin>>tmpIndex返回cin  (bool)cin?
    {
        
        if(tmpIndex >= 0 && tmpIndex <= 92)
            cout<<result[tmpIndex-1]<<"\n";
        cin.ignore();
    }
    return 0;
}


/*

https://leetcode-cn.com/problems/eight-queens-lcci/
class Solution {
public:
    vector<vector<string>>rval;
    void dfs(vector<string>& result, int startLineN, int maxLineN, vector<int>&colRecord)
    {
        if(startLineN == maxLineN)
        { 
            string s = "";
            int vecIndex = 0;
            for(int i = 0; i<maxLineN; i++)
            {//行
                int Qpos = colRecord[vecIndex++];
                for(int j = 0; j<maxLineN; j++)
                {//列
                    if(j == Qpos)
                    {
                        s.push_back('Q');
                    }else{
                        s.push_back('.');
                    }
                }
                result.push_back(s);
                s = "";
            }
            rval.push_back(result);
            result.clear();
            
            ;
        }else{
            for(int i = 0; i<maxLineN; i++)//将startLineN这一行摆到那一列上
            {
                //到这里来选列数
                if( find(colRecord.begin(), colRecord.end(), i) == colRecord.end() )
                {
                    int x = i+1, y = startLineN-1;
                    int lastIndex = colRecord.size()-1;
                    bool pointCanChose = true;
                    for(x,y; x<maxLineN&&y>=0; ++x, --y)
                    {
                        if(colRecord[lastIndex--] == x)
                        {
                            //重复点这个点不要;
                            pointCanChose = false;
                            break;
                        }
                    }
                    if(pointCanChose)
                    {
                        x = i-1;
                        y = startLineN-1;
                        lastIndex = colRecord.size()-1;
                        for(x,y; x>=0&&y>=0; --x, --y)
                        {
                            if(colRecord[lastIndex--] == x)
                            {
                                //重复点这个点不要;
                                pointCanChose = false;
                                break;
                            }
                        }
                    }
                    if(pointCanChose == false)
                    {
                        continue;
                    }else{
                        colRecord.push_back(i);
                        dfs(result, startLineN+1, maxLineN, colRecord);
                        colRecord.pop_back();
                    }
                }
            }
        }
    }


    vector<vector<string>> solveNQueens(int n) {
        vector<string>result;
        int startLineN = 0;
        vector<int> colRecord;
        int maxLineN = n;
        dfs(result, startLineN, maxLineN, colRecord);
        return rval;
    }
};


*/



