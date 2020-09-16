#include <iostream>
#include <vector>

using namespace std;

//质数又是素数，二者是相同的
int main()
{
    int n;
    cin>>n;
    vector<int> out(n+2,0);
    out[0] = 0;
    out[1] = 0;
    out[2] = 1;
    for(int i = 3; i<=n; i++)
    {
        out[i] = 1;
        for(int j = 2; j<i; j++)
        {
            if(i%j == 0)
            {
                out[i] = 0;
                break;
            }
        }
    }
    for(int i = 2; i<=n; i++)
        if(out[i])cout<<i<<" ";
    cout<<endl;
    return 0;
}



//https://www.nowcoder.com/test/question/done?tid=33968554&qid=105229#summary
//编程题:给定一个正整数，编程计算有多少对质数的和等于输入的正整数，输入值小于1000

#include <iostream>
#include <vector>
 
using namespace std;
 
 
//质数又是素数，二者是相同的
int main()
{
    int n;
    cin>>n;
    vector<int> out(n+2,0);
    out[0] = 0;
    out[1] = 0;
    out[2] = 1;
    for(int i = 3; i<=n; i++)
    {
        out[i] = 1;
        for(int j = 2; j<i; j++)
        {
            if(i%j == 0)
            {
                out[i] = 0;
                break;
            }
        }
    }
     
    int pair_count = 0;
    int first = 2;
    int last = n-1;
    while(first <= last)
    {
        if(!out[first])
        {
            first++;
            continue;
        }
         
        if(!out[last])
        {
            last--;
            continue;
        }
         
        int sum = first+last;
        if(sum == n)
        {
            pair_count++;
            first++;
            last--;
        }else if(sum > n)
        {
            last--;
        }else{
            first++;
        }
            
    }
    cout<<pair_count<<endl;
    return 0;
}
