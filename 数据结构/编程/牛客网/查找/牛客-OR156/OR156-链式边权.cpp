https://www.nowcoder.com/practice/dfb7f50ab15d479aafdb36ea934854cc?tpId=179&&tqId=34260&rp=1&ru=/activity/oj&qru=/ta/exam-other/question-ranking
#include <iostream>
#include <vector>
using namespace std;

int main()
{
    int n;   
    cin>>n;
    vector<int>in(n, 0);
    vector<int>record(n, 0);
    vector<int>out(n-1,0);
    for(int i = 0; i<n; i++)
        cin>>in[i];
    //对数组0到n开始遍历，每次的结果用到前一次的值
	//这个也是动态规划吗？虽说也是把整个计算问题，
	//分成了一小步一小步的计算
	
	//这个程序突出点：for循环中每一步做更多的事情
	//最外层循环有效利用前一次的值
    for(int i = 0; i<n; i++)
    {
        for(int j = i+1; j<n; j++)
        {
            if(in[i] != in[j])
            {
                record[i]++;
                record[j]--;
                //记录与j点匹配的所有的点，这样以来转移时就能由record[i-1]推出record[i]
                //就是这一步是最关键的
                //没次遍历的时候，在遍历中多记录一个信息，就越能减少运算量
            }
        }
        if(i>0)
            record[i] = record[i]+record[i-1];
        
        
    }
    
    for(int i = 0; i<n-1; i++)
    {
        cout<<record[i]<<" ";
    }
    
    
    return 0;
}

//下面这段代码是对上面代码的更进一步优化
//借助映射表，间O(n^2)的复杂度，优化成了O(n)
#include<bits/stdc++.h>
using namespace std;
int main(){
    int n;
    cin>>n;
    int a[n];
    long w[n];
    map<int,int> left,right;
    for(int i=0;i<n;i++){
        cin>>a[i];
        right[a[i]]++;
        left[a[i]]=0;
    }
    w[0]=0;
    for(int i=0;i<n-1;i++){
        right[a[i]]--;
        long t=w[i]-(i-left[a[i]])+(n-1-i-right[a[i]]);
        //w[i]-(i-left[a[i]])与上面的record[j--]起的功效一样
        //(n-1-i-right[a[i]])与上面的record[i]起的功效一样
        w[i+1]=t;
        left[a[i]]++;
    }
    for(int i=1;i<n;i++){
        if(i==n-1){
            cout<<w[i]<<endl;
        }else{
            cout<<w[i]<<' ';
        }
    }
    return 0;
}