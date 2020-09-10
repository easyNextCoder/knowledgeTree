#include <iostream>
#include <vector>
#include <random>

using namespace std;

int get_provit(vector<int>&vec, int first, int last)
{
    int mid = first+(last-first)/2;
    if(vec[first] > vec[last])
        swap(vec[first], vec[last]);
    if(vec[mid] < vec[first])
        swap(vec[mid], vec[first]);
    if(vec[mid] > vec[last])
        swap(vec[mid], vec[last]);

    return vec[mid];
}

void quick_sort(vector<int>&vec, int first, int last)
{

   if(last - first + 1 < 3)
   {
       if(first == last)return;
       else if(vec[first] > vec[last])
            swap(vec[first], vec[last]);

        return ;
    }   
    //下面的代码只能用来处理长度大于3的数组，原因如下：
    //1. j == last-1
    //2. --j;
    int provit = get_provit(vec, first, last);
    int mid = first+(last-first)/2;
    swap(vec[mid], vec[last-1]);
    int i = first;
    int j = last-1;
 
    for(;;)
    {
        while(vec[++i] < provit){;}
        while(vec[--j] > provit){;}
        if(i<j)
        {
            swap(vec[i], vec[j]);
        }else{
            break;
        }
    }
    cout<<i<<":"<<j<<endl;
    cout<<vec[last-1]<<endl;
    swap(vec[i], vec[last-1]);
    quick_sort(vec, first, i-1);
    quick_sort(vec, i+1, last);
}

int main()
{
    vector<int> vec = {9,8,7,6,5,4,3,2,1};
    default_random_engine e;
    int count = 1000;
    while(count--)
    {
        vec.push_back(e()%(10000));
    }
    cout<<"original array is:"<<endl;
    for(auto item:vec)
    {
        cout<<item<<" ";
    }
    quick_sort(vec, 0, vec.size()-1);
    cout<<"sorted array is:"<<endl;
    for(auto item:vec)
    {
        cout<<item<<" ";
    }
    cout<<endl;
    return 0;
}