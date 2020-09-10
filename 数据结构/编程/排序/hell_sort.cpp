#include <iostream>
#include <vector>
#include <algorithm>


using namespace std;

//选择排序，选择最小的放在前面
//插入排序，就是整体把整个队列向后移动，然后插入到合适的位置

void hell_sort(vector<int>& vec, int first, int last)
{
    for(int i = first; i<last; i++)
    {
        for(int N = (last-first)+1; N>=1; N/=2)
        for(int j = first+1; j<=last; j+=N)
        {
            if(vec[i]>vec[j])
                swap(vec[i], vec[j]);
        }
    }
}

int main()
{
    vector<int> in = {4,5,6,7,8,9,10};//{10, 9, 8,7,6,5,4,3,2,1,0};
    vector<int> in1 = {10, 9, 8,7,6,5,4,3,2,1,0};
    //sort(in.begin(), in.end());
    hell_sort(in1, 0, in1.size()-1);
    for(auto item:in1)
        cout<<item<<endl;

    return 0;
}
