//这个测试代码是由力扣：https://leetcode-cn.com/problems/rank-from-stream-lcci/
//引申出来的测试树状数组的测试代码
#include <iostream>
#include <list>
#include <random>
#include <string.h>

using namespace std;

const int LEN = 100;


int a[LEN];
int c[LEN];

int lowbit(int index)
{
    //cout<<"index:"<<index<<endl;
    //cout<<"index&(-index):"<<( index&(-index) )<<endl;
    return index&(-index);
}

int getSum(int index)
{
    int rval = 0;
    while(index > 0)
    {
        rval += c[index];
        index -= lowbit(index);
    }
    return rval;
}

void update(int index, int value)
{
    a[index] += value;
    while(index < LEN)
    {
        //cout<<index<<" ";
        c[index] += value;
        //index是在翻倍增加，将时间复杂度降到了O(nlogn)
        index += lowbit(index);
    }
}

int main()
{

    for(int i = 1; i<LEN; i++)
    {
        update(i, i);
    }

    cout<<"the result is:"<<endl;
    update(1, 100);
    for(int i = 1; i<LEN; i++)
    {
        cout<<getSum(i)<<" ";
    }

    return 0;
}