#include <iostream>
#include <vector>

using namespace std;

int f[101][101];
int sum[101];
class Solution {
public:
    int mergeStones(vector<int>& stones, int K) {
        int size = stones.size();
        for(int i = 1; i<=size; i++)
        {
            sum[i] = stones[i-1] + sum[i-1];
            //cout<<sum[i]<<endl;
        }

        
        int seqLen = K;
        for(int len = K; len<=size; len+=K-1)//先枚举区间长度
            for(int i = 1; i+len-1<=size; i++)//再枚举区间左端点
            {
                int j = i+len-1;////再枚举区间右端点
                f[i][j] = 1e8;
                for(int k = i; k+seqLen-1<=j; k++)  
                {   
                    cout<<"k is:"<<k<<endl;
                    if(len == K)
                    {
                        f[i][j] = sum[j] - sum[i-1];
                        cout<<"len = 3:"<<i<<j<<endl;
                    }else{
                        f[i][j] = min(f[i][j], f[k][k+len-(K)]+sum[j]-sum[i-1]);
                        cout<<"i i+len-K:"<<i<<" "<<i+len-K<<f[i][i+len-(K)]+sum[j]-sum[i-1]<<endl;
                    }
                    /*
                    if(k-i+1<seqLen)f[i][k] = sum[k]-sum[i-1];
                    if(j-(k+seqLen-1)+1<seqLen) f[k+seqLen-1][j] = sum[j]-sum[k+seqLen-1-1];
                    f[i][j] = min(f[i][j], f[i][k]+f[k+seqLen-1][j]+sum[j]-sum[i-1]);
                    cout<<"f[i][j]:"<<i<<" "<<j<<":"<<f[i][j]<<endl;
                    cout<<"f[i][k]:"<<i<<" "<<k<<":"<<f[i][k]<<endl;
                    cout<<"f[k+swqLen-1][j]:"<<k+seqLen-1<<" "<<j<<":"<<f[k+seqLen-1][j]<<endl;
                    */
                }   
            }
        return f[1][size];
    }
};

int main()
{
    Solution solution;
/*    vector<int> vec = {3,5,1,2,6};
    cout<<solution.mergeStones(vec, 3);
*/
    vector<int> vec = {3,2,4,1};
    cout<<solution.mergeStones(vec, 2);


    return 0;
}